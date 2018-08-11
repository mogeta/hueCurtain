package main

import (
	"github.com/spf13/viper"
	"fmt"
	"github.com/amimof/huego"
)

type Config struct {
	User struct {
		Name string `toml:"username"`
	}
}

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")               // 現在のワーキングディレクトリを探索することもできる
	err := viper.ReadInConfig()            // 設定ファイルを探索して読み取る
	if err != nil {                        // 設定ファイルの読み取りエラー対応
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}

	bridge, _ := huego.Discover()
	bridge.GetUsers()
	user := viper.GetStringMapString("User")["username"]
	bridge = bridge.Login(user)
	light, _ := bridge.GetLight(1)
	light.On()
}
