package main

import (
	"kosei-web/api"
)

func main() {
	api.Run()
}

// docker run -d -p 3306:3306 --name mysql-docker-container -e MYSQL_ROOT_PASSWORD=gotest -e MYSQL_DATABASE=gotest -e MYSQL_USER=sergey -e MYSQL_PASSWORD=gotest mysql:8.0.34
