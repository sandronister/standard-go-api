package main

import "github.com/sandronister/standart-go-api/configs"

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	print(configs.DBDriver)
}
