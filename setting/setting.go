package setting

import (
	"fmt"
	"io"
	"log"
	"os"
)
import "gopkg.in/yaml.v3"

type Config struct {
	Redis     Redis     `yaml:"redis"`
	Database  Database  `yaml:"database"`
	WebServer WebServer `yaml:"web_server"`
}

type Redis struct {
	host string `yaml:"host"`
	port int    `yaml:"port"`
	db   int    `yaml:"db"`
}

type Database struct {
	host     string `yaml:"host"`
	user     string `yaml:"user"`
	db       string `yaml:"db"`
	password string `yaml:"password"`
	port     int    `yaml:"port"`
}

type WebServer struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Prefork bool   `yaml:"Prefork"`
}

func Load_config() Config {
	err := copy_file("./dev.yaml", "./env.yaml")
	if err != nil {
		log.Fatal("copy file error")
	}
	content, err := os.ReadFile("./env.yaml")
	if err != nil {
		log.Fatal("read file error")
	}
	config := Config{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("unmarshal error")
	}
	fmt.Printf("config: %v", config)
	return config
}

func copy_file(src, dst string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("close file: %s error", file.Name())
		}
	}(file)
	dst_file, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func(dst_file *os.File) {
		err := dst_file.Close()
		if err != nil {
			log.Fatalf("close file: %s error", dst_file.Name())
		}
	}(dst_file)

	_, err = io.Copy(dst_file, file)
	if err != nil {
		return err
	}
	return nil
}
