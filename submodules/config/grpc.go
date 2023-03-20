package config

import "github.com/spf13/viper"

var GRPC struct {
	Port int
}

func init() {
	GRPC.Port = viper.GetInt("GRPC_PORT")
}
