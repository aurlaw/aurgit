package main

import (
	"encoding/json"
	"github.com/aurlaw/aurlog"
	"os"
)

// Represents the Server Application
type Application struct {
	Configuration // embedded
	Log           *aurlog.AurLog
}

// Represents the Server configuration
type Configuration struct {
	Port string
}

// Configures the application intance
func ConfigureApplication(configFile *string) *Application {

	conf := loadConfiguration(configFile)
	lc := aurlog.LogConfiguration{LogFile: "aurgit.log"}
	lc.IsInfo = true
	lc.IsError = true

	alog := aurlog.Configure(&lc)

	return &Application{conf, alog}
}

// loads json configuration
func loadConfiguration(configFile *string) Configuration {
	// log.Println("Load configuration file:", *configFile)

	file, _ := os.Open(*configFile)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}
