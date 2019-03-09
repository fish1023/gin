package setting

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Database 数据库配置
type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

// Redis redis配置
type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Pass string `yaml:"pass"`
}

// APISys api模块配置
type APISys struct {
	Port string `yaml:"port"`
}

// App 系统模块配置
type App struct {
	API APISys
}

type Common struct {
	LogPath string `yaml:"logpath"`
}

//Config   系统配置配置
type Config struct {
	Database map[string]Database `yaml:"Database"`
	Redis    map[string]Redis    `yaml:"Redis"`
	App      App                 `yaml:"App"`
	Common   Common              `yaml:"Common"`
}

// Setting is
var Setting = &Config{}

// init is init config
func init() {
	config, err := ioutil.ReadFile("conf/app.yaml")
    if err != nil {
		fmt.Print(err)
	}

	yaml.Unmarshal(config, &Setting)
}