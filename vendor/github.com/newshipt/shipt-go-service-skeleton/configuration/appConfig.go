package configuration

import (
	"github.com/newshipt/shipt-go-utils/configuration"
)

//AppConfiguration struct
type AppConfiguration struct {
	configuration.CommonConfig
}

//GetAppConfig app server configuration
func GetAppConfig() AppConfiguration {
	configState := configuration.NewConfigState()

	/* Sample of how this will be extended as we start building out this service

	AppSpecificSetting := os.Getenv("APP_SPECIFIC_SETTING")

	if AppSpecificSetting == "" {
		configState.SetMissing("APP_SPECIFIC_SETTING")
	}

	*/
	appConfig := AppConfiguration{
		CommonConfig: configuration.NewCommonConfig(),
	}

	configState.Validate()

	return appConfig
}
