package main

import (
	"github.com/spf13/viper"
	"fmt"
	"github.com/amimof/huego"
	"runtime"
)

type Config struct {
	User struct {
		Username string `toml:"username"`
	} `toml:"User"`
	Env struct {
		Start    string `toml:"start"`
		End      string `toml:"end"`
		Interval int    `toml:"interval"`
		Inc      int    `toml:"inc"`
	} `toml:"Env"`
}

type HueManager struct {
	Username         string
	Briinc           int
	ColorTemperature uint16
}

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}

	manager := HueManager{c.User.Username, c.Env.Inc, 230}
	execute(c.Env.Start, c.Env.End, c.Env.Interval, manager.fade)
	runtime.Goexit()
}

func (h HueManager) fade() {
	bridge, _ := huego.Discover()
	bridge.GetUsers()
	bridge = bridge.Login(h.Username)
	group, _ := bridge.GetGroup(1)
	s := huego.State{On: true, BriInc: h.Briinc, Ct: h.ColorTemperature}
	group.SetState(s)

}
