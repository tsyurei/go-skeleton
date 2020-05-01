package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "go-skeleton",
	Short: "basic skeleton template",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var configFile string

func init() {
	cobra.OnInitialize(InitConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Config File (Default Config File is config.yaml)")
}

func InitConfig() {
	viper.SetConfigType("yaml")

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		_, b, _, _ := runtime.Caller(0)
		basepath := filepath.Dir(b)
		ix := strings.LastIndex(basepath, "/");

		viper.AddConfigPath(".")
		viper.AddConfigPath(basepath[0:ix])
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("unable to find config %v \n", err)
		os.Exit(1)
	}
}
