package configs

import (
	"encoding/json"
	"os"
)

// Config 整个项目的配置
type Config struct {
	Mode        string `json:"mode"`
	Port        int    `json:"port"`
	*LogConfig  `json:"log"`
	Neo4jConfig `json:"neo4JConfig"`
}

type Neo4jConfig struct {
	Url          string `json:"url"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"databaseName"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

// Conf 全局配置变量
var Conf = new(Config)

// Init 初始化配置；从指定文件加载配置文件
func Init(filePath string) error {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, Conf)
}
