package setting

import (
	"fmt"
	"go-hello/web/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Db   int    `yaml:"db"`
}

type Database struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Db       string `yaml:"db"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

type WebServer struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Prefork bool   `yaml:"Prefork"`
}

func Load_config() Config {
	err := copyFile("./dev.yaml", "./env.yaml")
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
		log.Fatal("unmarshal error", err.Error())
	}
	fmt.Printf("config: %+v", config)
	return config
}

func copyFile(src, dst string) error {
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

var Db *gorm.DB

func Connect(config Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Db)
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	Db.AutoMigrate(&models.Order{})
	return nil
}
