package main

import (
	"fmt"
	"os"

	"go_redis/config"
	"go_redis/lib/logger"
	// RedisServer "go_redis/redis/server"
	// "go_redis/tcp"
)

var banner = `
   ______                       ___
  / ____/___      ___ ___  ____/ (_)____
 / / __/ __ \    / __/ _ \/ __  / / ___/
/ /_/ / /_/ /__ / / |  __/ /_/ / (__  )
\____/\____/__ /_/   \___\__,_/_/____/
`

var defaultProperties = &config.ServerProperties{
	Bind:           "0.0.0.0",
	Port:           6399,
	AppendOnly:     false,
	AppendFilename: "",
	MaxClients:     1000,
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

func main() {
	fmt.Print(banner)
	logger.Setup(&logger.Settings{
		Path:       "logs",
		Name:       "godis",
		Ext:        "log",
		TimeFormat: "2006-01-02",
	})
	// export CONFIG="xxx"
	configFilename := os.Getenv("CONFIG")
	if configFilename == "" {
		if fileExists("redis.conf") {
			fmt.Println("Loading config file: redis.conf ...")
			config.SetupConfig("redis.conf")
		} else {
			fmt.Println("Using default properties.")
			config.Properties = defaultProperties
		}
	} else {
		fmt.Println("Loading config file: ", configFilename, " ...")
		config.SetupConfig(configFilename)
	}
	fmt.Println("Current configuration: ", config.Properties.String())

	// err := tcp.ListenAndServeWithSignal(&tcp.Config{
	// 	Address: fmt.Sprintf("%s:%d", config.Properties.Bind, config.Properties.Port),
	// }, RedisServer.MakeHandler())
	// if err != nil {
	// 	logger.Error(err)
	// }
}
