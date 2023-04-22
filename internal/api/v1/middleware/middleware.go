package middleware

/*
func XApiKeyMiddleware(context *fiber.Ctx) error {
	requestApiKey := context.Get("XAPIKEY")
	serverApiKey := utils.GetEnvVariable("XAPIKEY")
	fmt.Println("requestApiKey")
	fmt.Println(requestApiKey)
	fmt.Println("serverApiKey")
	fmt.Println(serverApiKey)
	// context.h

	if requestApiKey != serverApiKey {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "unauthorized",
		})
	}

	return context.Next()
}
*/
