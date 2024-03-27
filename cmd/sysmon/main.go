// command sysmon collects and sends several parameters to clients
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/Elyas-Crimean/sysmon/internal/collector"
	"github.com/Elyas-Crimean/sysmon/internal/server"
	"github.com/Elyas-Crimean/sysmon/internal/storage"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	configFile := pflag.StringP("config", "c", "", "файл конфигурации")
	pflag.Parse()
	if *configFile != "" {
		viper.SetConfigFile(*configFile)
	}
	viper.SetConfigName("sysmon")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.SetDefault("historyhours", 24)
	viper.SetDefault("listenport", 8000)
	err := viper.ReadInConfig()
	if err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			fmt.Println("Config file not found")
			return
		}
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	st := storage.NewStorage(time.Hour * time.Duration(viper.GetInt("historyhours")))
	c := collector.NewCollector(st)
	c.Run()
	srv := server.New(viper.GetInt("listenport"))
	if srv == nil {
		return
	}
	srv.Run(st)
}
