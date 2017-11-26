package configuration

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// CommonConfig ...
type CommonConfig struct {
	Server             string
	Environment        string
	RollbarToken       string
	AllowedOrigins     []string
	NewRelicLicenseKey string
	NewRelicAppName    string
	DocumentsFolder    string
	GraphiteHost       string
	GraphiteAPIKey     string
	AuthURL            string
}

// NewCommonConfig configuration setup
func NewCommonConfig(appName string) CommonConfig {
	configState := NewConfigState()

	// Forcing Viper to fall back on ENV if value is not in config file.
	// Note: Do not call viper.SetEnvPrefix, which will set prefix for all Env vars and break things
	viper.AutomaticEnv()

	// Set config file name: {appName}.json
	viper.SetConfigName("." + appName)
	viper.SetConfigType("json")

	// Tell Viper where to look for config
	viper.AddConfigPath("$HOME/")         // home dir
	viper.AddConfigPath("$HOME/.config/") // config dir
	viper.AddConfigPath("./")             // project root
	viper.AddConfigPath("../")            // for packages

	// Load up config
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("Didn't find config file; falling back on ENV")
	}

	Server := viper.GetString("server_host")
	port := viper.GetString("port")

	if port != "" {
		Server = Server + ":" + port
	}

	config := CommonConfig{
		Server:             Server,
		Environment:        viper.GetString("environment"),
		RollbarToken:       viper.GetString("rollbar_token"),
		AllowedOrigins:     strings.Split(viper.GetString("allowed_origins"), ";"),
		NewRelicLicenseKey: viper.GetString("new_relic_license_key"),
		NewRelicAppName:    viper.GetString("new_relic_application_name"),
		DocumentsFolder:    viper.GetString("documents_folder"),
		GraphiteHost:       viper.GetString("hosted_graphite_host"),
		GraphiteAPIKey:     viper.GetString("hosted_graphite_api_key"),
		AuthURL:            viper.GetString("auth_url"),
	}

	if config.Environment == "production" {
		if config.NewRelicLicenseKey == "" {
			configState.SetMissing("NEW_RELIC_LICENSE_KEY")
		}

		if config.NewRelicAppName == "" {
			configState.SetMissing("NEW_RELIC_APPLICATION_NAME")
		}

		if config.RollbarToken == "" {
			configState.SetMissing("ROLLBAR_TOKEN")
		}

		if config.GraphiteHost == "" {
			configState.SetMissing("HOSTED_GRAPHITE_HOST")
		}

		if config.GraphiteAPIKey == "" {
			configState.SetMissing("HOSTED_GRAPHITE_API_KEY")
		}
	}

	configState.Validate()

	return config
}
