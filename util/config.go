package util

import (
	"ego/iface"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	TcpServer      iface.IServer // server
	Host           string        // ip
	Port           int           // 端口
	Name           string        // 服务名称
	Version        string        // 版本
	MaxConn        int           // 最大连接数
	MaxPackageSize uint32        // 数据包大小
}

func (c *Config) Reload() {
	data, err := ioutil.ReadFile("app.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &Config{}); err != nil {
		panic(err)
	}
}

var Configure *Config

func init() {
	Configure = &Config{
		Host:           "0.0.0.0",
		Port:           10000,
		Name:           "Ego",
		Version:        "v1.0.0",
		MaxConn:        1000,
		MaxPackageSize: 1048576,
	}
	Configure.Reload()
}
