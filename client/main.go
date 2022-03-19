package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main(){
	initConfig()

	app := fiber.New()
	app.Get("/", Send)

	app.Listen(":4000")


}

func Send(c *fiber.Ctx) error {
	u := viper.Get("url")
	uri,_ := u.(string)
	resp,err := http.Get(uri)

	if err != nil{
		return c.SendString("htt.Get:" + err.Error() + "--uri--"+uri)
	}else{
		defer resp.Body.Close()
		bt,err := ioutil.ReadAll(resp.Body)
		if err != nil{
			return c.SendString("ReadAll" + err.Error() + "--uri--"+uri)
		}
		return c.SendString("htt.Get.res" + string(bt) + "--uri--"+uri)
	}
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

