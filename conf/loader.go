package conf

import (
	"fmt"
	"github.com/itzmeerkat/mentally-friendly-infra/config"
)

type MySqlConfig struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
	Address string `yaml:"Address"`
}
func (m *MySqlConfig) GetMasterDSN() string {
	return fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", m.Username, m.Password, m.Address, m.Database)
}

type EnvConfig struct {
	DEALMAKER_ENV string
}
func (e *EnvConfig) IsProd() bool {
	return e.DEALMAKER_ENV == "prod"
}

type SendGridConfig struct {
	SENDGRID_API_KEY string
}

type AllConfig struct {
	MySqlConfig
	SendGridConfig
	EnvConfig
}

var Conf AllConfig
func init() {
	config.LoadEnvVar(&Conf.EnvConfig)
	if Conf.EnvConfig.DEALMAKER_ENV == "" {
		panic("Set environment before launch")
	}
	confPath := "./conf/" + Conf.EnvConfig.DEALMAKER_ENV +".yml"
	config.LoadConfigFile(confPath, &Conf.MySqlConfig)
	config.LoadEnvVar(&Conf.SendGridConfig)
	//fmt.Println(conf)
}
