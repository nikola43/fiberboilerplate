Make sure you have docker and docker-compose installed. Refer this official documentation for installation

mac
docker pull --platform linux/x86_64 mysql

//build our custom image(i.e dockerfile provided on build key)
docker-compose build 

//starts all our container(service) configured on yaml file
docker-compose up

docker-compose build && docker-compose up