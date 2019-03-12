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

// RedisServer redis配置
type RedisServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Pass string `yaml:"pass"`
}

// Redis
type RedisIDC struct {
    Master map[string]string   `yaml:"master"`
    Slave  []map               `yaml:"master"`
}

// App 系统模块配置
type App struct {
	Port string `yaml:"port"`
    LogPath string `yaml:"logPath"`
}

//Config   系统配置配置
type Config struct {
	Database map[string]Database `yaml:"Database"`
	Redis    map[string]RedisIDC `yaml:"Redis"`
	App      map[string]App      `yaml:"App"`
}

// Setting is
var Setting = Config{}

// init is init config
func init() {
	config, err := ioutil.ReadFile("conf/app.yaml")
    if err != nil {
		fmt.Print(err)
	}

	yaml.Unmarshal(config, &Setting)
}
