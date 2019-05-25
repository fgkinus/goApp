//this package the global configuration for logging
package config

import (
	"github.com/sirupsen/logrus"
)

func init() {
	Logger.SetLevel(logrus.InfoLevel)
	//If you wish to add the calling method as a field, instruct the logger via
	Logger.SetReportCaller(true)
}

var Logger = logrus.New()
