package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Ethereum struct {
	Network string `toml:"network"`
	ChainId int `toml:"chainId"`
	KeyStore string `toml:"keyStore"`
	ContractAddress string `toml:"contractAddress"`
	UsdtAddress string `toml:"usdtAddress"`
	MubsAddress string `toml:"mubsAddress"`
	UsdtsAddress string `toml:"usdtsAddress"`
	UsdtReceiver string `toml:"usdtReceiver"`
}

type Database struct {
	DBHost       string  `toml:"db_host"`
	DBPort       string  `toml:"db_port"`
	DBSchema     string  `toml:"db_schema"`
	DBUserName 	 string  `toml:"db_username"`
	DBPassword   string  `toml:"db_password"`
	DBArgs       string  `toml:"db_args"`
}
type Redis struct {
	RedisNetwork string `toml:"redis_network"`
	RedisPassword string `toml:"redis_password"`
}


type configuration struct {
	Database Database `toml:"database"`
	Ethereum Ethereum `toml:"ethereum"`
	Redis Redis `toml:"redis"`
}
var Cfg  *configuration
func Config() *configuration {
	if _, err := toml.DecodeFile("./config/config.toml", &Cfg); err != nil {
		log.Println(err)
	}
	return Cfg
}
