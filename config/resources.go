package config

import "github.com/spf13/viper"

var ResourcesFilePath string

func setupResourcesConfig() {
	ResourcesFilePath = viper.GetString("resources_file_path")
}
