package main

import (
	"kreditplus-api/config"
	"kreditplus-api/container"
)

func main() {
	config.InitConfiguration()
	containerApp := container.InitContainer()
	container.InitRouter(containerApp)
}
