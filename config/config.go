package config

import (
	"fmt"
	"time"
)

type Config struct {
	TemplatePath string
}

var ServerConfig Config

func InitConfig(templatePath string) {
	if templatePath == "" {
		fmt.Println("Template Path is null")
	}
	ServerConfig = Config{TemplatePath: templatePath}
}

func execute() {
	for {
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
