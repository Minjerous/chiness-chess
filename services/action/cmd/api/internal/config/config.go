package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type UserConfig struct {
	Name     string     `yaml:"Name"`
	MysqlCfg MysqlCfg   `yaml:"Mysql"`
	Redis    RedisCfg   `yaml:"Redis"`
	Host     string     `yaml:"Host"`
	UserRpc  UserRpcCfg `yaml:"UserRpc"`
	Jwt      Jwt        `yaml:"Jwt"`
	Email    Email      `yaml:"Email"`
	Mq       MqCfg      `yaml:"Mq"`
}

//用户信息缓存
type RedisCfg struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
}

type MysqlCfg struct {
	DataSource string `yaml:"DataSource"`
}
type UserRpcCfg struct {
	Hosts string `yaml:"Host"`
	Key   string `yaml:"Key"`
}

type Jwt struct {
	RefreshSecret string `yaml:"RefreshSecret"`
	AccessSecret  string `yaml:"AccessSecret"`
}

type Email struct {
	ServiceEmail string `yaml:"ServiceEmail"`
	ServicePwd   string `yaml:"ServicePwd"`
	SmtpPort     string `yaml:"SmtpPort"`
	SmtpHost     string `yaml:"SmtpHost"`
}

type MqCfg struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	PassWord string `yaml:"PassWord"`
}

var UserCfg *UserConfig

func GetUserCfg() *UserConfig {
	return UserCfg
}

func PareConfig() {
	file, err := os.Open("./cmd/api/internal/config/user.yml")
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bytes, &UserCfg)
	if err != nil {
		panic(err)
	}
}
