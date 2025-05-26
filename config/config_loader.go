package config

import (
	_ "embed"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DB    DBConfig
	S     ServerConfig
	JWT   JWTConfig
	MQTT  MQTTConfig
	Redis RedisConfig
}

func LoadConfigFile(path string) (*Config, error) {
	var conf Config
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	//DBConfig

	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, fmt.Errorf("DB_HOST is not set in the environment")
	}

	portStr := os.Getenv("DB_PORT")
	if portStr == "" {
		return nil, fmt.Errorf("DB_PORT is not set in the environment")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("Invalid DB_PORT value: %v", err)
	}

	username := os.Getenv("DB_USER")
	if username == "" {
		return nil, fmt.Errorf("DB_USER is not set in the environment")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is not set in the environment")
	}

	database := os.Getenv("DB_NAME")
	if database == "" {
		return nil, fmt.Errorf("DB_NAME is not set in the environment")
	}

	conf.DB = DBConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
	}

	//Server Config

	host = os.Getenv("HOST")
	port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	conf.S = ServerConfig{
		Host: host,
		Port: port,
	}

	//JWT Config

	key := os.Getenv("JWT_KEY")
	if key == "" {
		return nil, fmt.Errorf("JWT_KEY is not set in the environment")
	}
	exp, err := strconv.Atoi(os.Getenv("JWT_EXP"))
	if err != nil {
		return nil, err
	}
	conf.JWT = JWTConfig{
		Key:    key,
		Expire: time.Duration(exp) * time.Minute,
	}

	//MQTT Config

	mqlogin := os.Getenv("MQTT_LOGIN")
	if mqlogin == "" {
		return nil, fmt.Errorf("MQTT_LOGIN is not set in the environment")
	}
	mqpass := os.Getenv("MQTT_PASSWORD")
	mqhost := os.Getenv("MQTT_HOST")
	if mqhost == "" {
		return nil, fmt.Errorf("DB_HOST is not set in the environment")
	}
	mqport := os.Getenv("MQTT_PORT")
	conf.MQTT = MQTTConfig{
		Host:     mqhost,
		Port:     mqport,
		Login:    mqlogin,
		Password: mqpass,
	}

	// Redis config

	addr := os.Getenv("REDIS_ADDR")
	database = os.Getenv("REDIS_DB")
	databaseInt, err := strconv.Atoi(database)
	if err != nil {
		return nil, err
	}

	conf.Redis = RedisConfig{
		Address:  addr,
		Database: databaseInt,
	}

	return &conf, nil
}
