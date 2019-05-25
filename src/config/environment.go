/*
This package implements application configuration.
The methods defined are used to read compile time env variables and read time variables
*/

package config

import (
	"flag"
	"os"

	"github.com/jinzhu/configor"

	"github.com/joho/godotenv"
)

type CMDArgs struct {
	ConfigFilePath string
}

type Config struct {
	Database struct {
		Uri string
	}
	APPName string
}

var Configuration = Config{}

func GetConfig() Config {
	return Configuration
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getCurrentWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		Logger.Error(err)
	}
	return dir
}

// ReadCompileTimeEnv is supposed to read compile time env variables
func ReadCompileTimeEnv() {
	dir := getCurrentWorkingDirectory()
	err := godotenv.Load(dir + "/src/config/.env")
	if err != nil {
		Logger.Fatal(err)
	}
}

// ReadRuntimeConfig is supposed to read the env variables required during runtime including db setup vars ,
// initial credentials etc
func ReadRuntimeConfig(path string) {
	err := configor.Load(&Configuration, path)
	if err != nil {
		Logger.Fatal(err)
	}
	Logger.Debug(Configuration)
}

// GetCMDArgs is a method that is used to read command line arguments including the config file for app initialisation
// the config file defaults to the current directory with the config file
func GetCMDArgs() CMDArgs {
	arguments := CMDArgs{}
	dir := getCurrentWorkingDirectory()
	configFileUrl := flag.String("config", dir+"/config.yml", "The path to the configuration file for the application")
	flag.Parse()
	arguments.ConfigFilePath = *configFileUrl
	return arguments
}
