package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "EssayGenie",
	Short: "EssayGenie is a web application that writes a essay in English by just putting keyword in Korean",
	Long:  "EssayGenie is a web application that writes a essay in English by just putting keyword in Korean.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	profile := initProfile()
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(profile)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initProfile() string {
	profile := os.Getenv("go_profile")
	if profile == "" {
		profile = "dev"
	}
	fmt.Println("profile:", profile)
	return profile
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(serveCmd)
}
