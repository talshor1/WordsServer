package config

import (
	"fmt"
	"github.com/spf13/viper"
	"regexp"
	"strconv"
)

type Config struct {
	PerSec  int
	PerMin  int
	Address string
	Port    string
}

func ReadAndValidateConfig(fname, ftype, fpath string) (error, *Config) {
	var (
		serverAddress string
		serverPort    string
		perMin        int
		perSec        int
	)

	viper.SetConfigName(fname)
	viper.SetConfigType(ftype)
	viper.AddConfigPath(fpath)

	err := viper.ReadInConfig()
	if err != nil {
		return err, nil
	}

	if !viper.IsSet("server.address") || !viper.IsSet("server.port") {
		return fmt.Errorf("missing required fields in server configuration"), nil
	}

	if !isValidAddress(viper.GetString("server.address")) {
		return fmt.Errorf("server addrress is not valid"), nil
	}

	if !isValidPort(viper.GetString("server.port")) {
		return fmt.Errorf("server port is not valid"), nil
	}

	if !viper.IsSet("limiter.allowed_per_sec") || !viper.IsSet("limiter.allowed_per_min") {
		return fmt.Errorf("missing required fields in limiter configuration"), nil
	}

	if err, perSec, perMin = isValidLimiter(); err != nil {
		return fmt.Errorf("%v", err), nil
	}

	serverAddress = viper.GetString("server.address")
	serverPort = viper.GetString("server.port")

	c := &Config{
		Address: serverAddress,
		Port:    serverPort,
		PerSec:  perSec,
		PerMin:  perMin,
	}
	return nil, c
}

func isValidLimiter() (error, int, int) {
	min := 1
	max := 20
	perSec := viper.GetString("limiter.allowed_per_sec")
	perMin := viper.GetString("limiter.allowed_per_min")

	perSecValue, err := strconv.Atoi(perSec)
	if err != nil {
		return fmt.Errorf("limiter perSecValue config is not valid"), 0, 0
	}

	if perSecValue < min || perSecValue > max {
		return fmt.Errorf("limiter perSecValue config is not valid"), 0, 0
	}

	perMinValue, err := strconv.Atoi(perMin)
	if err != nil {
		return fmt.Errorf("limiter perMinValue config is not valid"), 0, 0
	}

	if perMinValue < min || perMinValue > max {
		return fmt.Errorf("limiter perMinValue config is not valid"), 0, 0
	}
	return nil, perSecValue, perMinValue
}

func isValidAddress(address string) bool {
	return regexp.MustCompile(`^([0-9a-fA-F:.]+|\d+\.\d+\.\d+\.\d+)$`).MatchString(address)
}

func isValidPort(port string) bool {
	portNum, err := strconv.Atoi(port)
	return err == nil && portNum >= 1 && portNum <= 65535
}
