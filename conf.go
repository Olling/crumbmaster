package main

import (
	"flag"
	"github.com/olling/slog"
)

type CronMessage struct {
	Cron string
	Message Message
}

type Message struct {
	To User
	Subject string
	Message string
}

type MessageData struct {
	Accepted bool
	Reposible User
	Fallback User
	NextDate string
}

type Configuration struct {
	ConfigurationType string
	PathJsonConfiguration string
	HttpPort int
	HttpTlsCert string
	HttpTlsKey string
	HttpTlsPort int
	Debug bool
	Timeout int
	Tokens []string
	CronBreadDate string
	CronRemind string
	CronNext string
	CronMessages []CronMessage
}

var (
	ConfigurationPath = flag.String("configurationpath","/etc/crumbmaster.conf","(Optional) The path to the configuration file")
	CurrentConfiguration Configuration
)

func InitializeConfiguration() {
	slog.PrintDebug("Initializing configuration")
	flag.Parse()
	slog.PrintDebug("Configuration Path: " + *ConfigurationPath)
	ReadJsonFile(*ConfigurationPath,&CurrentConfiguration)
}
