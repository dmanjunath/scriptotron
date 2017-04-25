package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var config tomlConfig

type tomlConfig struct {
	Command   				  string
	Logfile   				  string
	ExcludeErrorKeywords      []string
	IncludeErrorKeywords	  []string
	Actions					  []string
	Gmail 					  gmail
	Twilio 					  twilio
}

type gmail struct {
	Email 	  string
	Password  string
}

type twilio struct {
	AccountSid  	string
	AuthToken   	string
	ToPhoneNumber 	string
	FromPhoneNumber string
}

func parseConfig(){
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
}