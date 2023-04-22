package utils

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/fiberboilerplate/internal/api/v1/models"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateClientToken(pk *rsa.PrivateKey, email string, userId uint) (string, error) {
	claims := jwt.MapClaims{
		"id":    userId,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	fmt.Println(claims["id"])
	fmt.Println(claims["email"])
	fmt.Println(claims["exp"])

	// Generate encoded token and send it as response.
	// Create token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Generate encoded token and send it as response.
	t, err := jwtToken.SignedString(pk)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return "", err
	}
	return t, nil
}

func GetClientTokenClaims(context *fiber.Ctx) (*models.UserTokenClaims, error) {
	user := context.Locals("user").(*jwt.Token)
	if claims, ok := user.Claims.(jwt.MapClaims); ok && user.Valid {
		clientTokenClaims := &models.UserTokenClaims{}

		if claims["id"] != nil {
			clientTokenClaims.ID = uint(math.Round(claims["id"].(float64)))
		}

		if claims["email"] != nil {
			clientTokenClaims.Email = claims["email"].(string)
		}

		if claims["exp"] != nil {
			clientTokenClaims.Exp = uint(math.Round(claims["exp"].(float64)))
		}

		return clientTokenClaims, nil
	} else {
		return nil, errors.New("invalid claims")
	}
}

func GenerateRSAPrivateKey() error {
	rng := rand.Reader
	privateKey, err := rsa.GenerateKey(rng, 2048)
	if err != nil {
		return err
	}

	// Save private key to file
	privateFile, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	defer privateFile.Close()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	err = pem.Encode(privateFile, privateBlock)
	if err != nil {
		return err
	}

	return nil
}

func ReadRsaPrivateKeyFromFile() (*rsa.PrivateKey, error) {
	privateKeyFile, err := os.Open("private.pem")
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}
	defer privateKeyFile.Close()

	privateKeyFileinfo, _ := privateKeyFile.Stat()
	var size = privateKeyFileinfo.Size()
	privateKeyBytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	data, _ := pem.Decode([]byte(privateKeyBytes))
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKeyImported, nil
}
