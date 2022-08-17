package test

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

type testGlobalCfg struct {
	TcpServer      string `yaml:"tcp_server"`     // server对象
	IP             string `yaml:"ip"`             // 监听的ip
	TcpPort        int    `yaml:"tcp_port"`       // 监听的端口
	Name           string `yaml:"name"`           // Server name
	Version        string `yaml:"version"`        // server version
	MaxConn        int    `yaml:"maxConn"`        // 最大的客户端并发数目
	MaxPackageSize uint32 `yaml:"maxPackageSize"` // 读取包的长度
}

// server config
type TestConfig struct {
	GlobalConfig testGlobalCfg `yaml:"globalCfg"`
}

// init read error
func TestReadYaml(t *testing.T) {
	data, err := ioutil.ReadFile("../conf/config.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println("--------------------------")
	cfg := &TestConfig{}
	err = yaml.Unmarshal([]byte(data), cfg)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.GlobalConfig.Name != "GameServer" {
		t.Error("cfg err")
	}
	if cfg.GlobalConfig.MaxConn != 1000 {
		t.Error("err")
	}
	if cfg.GlobalConfig.MaxPackageSize != 512 {
		t.Error("err")
	}
	if cfg.GlobalConfig.TcpPort != 8888 {
		t.Error("err")
	}
}
