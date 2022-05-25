package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	Dburl         string `mapstructure:"DB_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/env")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		log.Fatalln("unable to read config file")
		return
	}

	err = viper.Unmarshal(&config)

	return
}

// func main() {
// 	a, err := LoadConfig()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(a.Port)
// }
