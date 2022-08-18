package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"leexsh/TCPGame/TCPGameServer/iface"
)

type globalCfg struct {
	TcpServer      iface.IServer `yaml:"tcp_server"`     // server对象
	IP             string        `yaml:"ip"`             // 监听的ip
	TcpPort        int           `yaml:"tcp_port"`       // 监听的端口
	Name           string        `yaml:"name"`           // Server name
	Version        string        `yaml:"version"`        // server version
	MaxConn        int           `yaml:"maxConn"`        // 最大的客户端并发数目
	MaxPackageSize uint32        `yaml:"maxPackageSize"` // 读取包的长度
}

// server config
type Config struct {
	GlobalConfig globalCfg `yaml:"globalCfg"`
}

func (g *Config) ReadConfig() {
	data, err := ioutil.ReadFile("conf/config.yml")
	fmt.Println(string(data))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &YmlConfig)
	if err != nil {
		panic(err)
	}
}

var YmlConfig *Config

func init() {
	// 未加载配置文件的时候
	YmlConfig = &Config{
		globalCfg{
			TcpServer:      nil,
			IP:             "127.0.0.1",
			TcpPort:        8888,
			Name:           "GameServer",
			Version:        "v1.0",
			MaxConn:        1000,
			MaxPackageSize: 512,
		}}
	YmlConfig.ReadConfig()
}
