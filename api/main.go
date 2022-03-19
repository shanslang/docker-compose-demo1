package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		str := viper.Get("name")
		name,_ := str.(string)
		return c.SendString("Hello, World 👋!【"+name)
	})

	app.Listen(":3000")
}

func initConfig(){
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}