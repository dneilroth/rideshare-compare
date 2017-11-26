package configuration

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

//ConfigState is a struct containing the needed data for
//telling us if the server has been configured correctly
type ConfigState struct {
	Messages []string
}

//NewConfigState is a constructor for a ConfigState struct
func NewConfigState() *ConfigState {
	return &ConfigState{
		Messages: []string{},
	}
}

//Getenv takes an env var name and tries to
//get that environment variable, returning an empty
//string if it can't. Also, it stores information about
//the env vars it can't find.
func (c *ConfigState) Getenv(envVar string) string {
	value := viper.GetString(envVar)
	if value == "" {
		c.SetMissing(envVar)
		return ""
	}

	return value
}

//Getint takes an env var name and tries to
//get that environment variable, returning 0
//if it can't. Also, it stores information about
//the env vars it can't find.
func (c *ConfigState) Getint(envVar string) int {
	value := c.Getenv(envVar)

	if value != "" {
		i, err := strconv.Atoi(value)

		if err != nil {
			c.SetMissing(envVar)
			return 0
		}
		return i
	}
	return 0
}

//SetMissing Adds a message to the Messages slice
//saying that a specific env var is missing
func (c *ConfigState) SetMissing(envVar string) {
	c.Messages = append(c.Messages, fmt.Sprintf("Please set the %s env var", envVar))
}

//Validate checks if there are any messages, and if
//there are then it logs them and os.Exit(1).
func (c *ConfigState) Validate() {
	if len(c.Messages) > 0 {
		for _, message := range c.Messages {
			log.Println(message)
		}
		os.Exit(1)
	}
}
