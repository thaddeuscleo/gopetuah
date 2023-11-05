package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thaddeuscleo/gopetuah/cmd/proxy"
)

var rootCmd = &cobra.Command{
	Use:   "gopetuah",
	Short: "Command Server for Route TCP packages to multiple upstream servers",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initConfig()
	rootCmd.AddCommand(proxy.ProxyCmd)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("fatal error config file: ", err)
	}
}
