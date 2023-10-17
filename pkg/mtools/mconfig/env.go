package mconfig

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	ApiKeySendgrid    string `mapstructure:"SENDGRID_APIKEY"`
	FromEmailSendgrid string `mapstructure:"SENDGRID_FROM_EMAIL"`
	UserNameSendgrid  string `mapstructure:"SENDGRID_USER_NAME"`
	ApiKeyMailGun     string `mapstructure:"MAILGUN_APIKEY"`
	Domain            string `mapstructure:"MAILGUN_DOMAIN"`
	Sender            string `mapstructure:"MAILGUN_SENDER"`
}

var GlobalEnv *Env

func NewEnv() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic("Error the fields are not valid")
	}

	viper.SetDefault("TIMEZONE", "UTC")

	err = viper.Unmarshal(&GlobalEnv)
	if err != nil {
		panic("environment cant be loaded")
	}
}
