package main

import (
	"github.com/amimof/huego"
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	User struct {
		Name string `toml:"username"`
	}
}

func main() {

	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	bridge, _ := huego.Discover()
	bridge.GetUsers()
	user := config.User.Name
	bridge = bridge.Login(user)
	light, _ := bridge.GetLight(1)
	light.On()
}
