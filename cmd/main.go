package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type param struct {
	name      string
	shorthand string
	value     interface{}
	usage     string
	required  bool
}

var (
	Version    string
	Build      string
	rootParams = []param{
		{name: "json-log", shorthand: "", value: false, usage: "output logs in json format"},
		{name: "verbose", shorthand: "v", value: false, usage: "enable verbose logs"},
	}
)

func setParams(params []param, command *cobra.Command) {
	for _, param := range params {
		switch v := param.value.(type) {
		case int:
			command.PersistentFlags().IntP(param.name, param.shorthand, v, param.usage)
		case string:
			command.PersistentFlags().StringP(param.name, param.shorthand, v, param.usage)
		case bool:
			command.PersistentFlags().BoolP(param.name, param.shorthand, v, param.usage)
		}
		if err := viper.BindPFlag(param.name, command.PersistentFlags().Lookup(param.name)); err != nil {
			panic(err)
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "envoy-authz",
	Short: "envoy-authz",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("running envoy-authz")
	},
}

func setupLogging() {

	if viper.GetBool("verbose") {
		log.SetLevel(log.DebugLevel)
	}
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		DisableColors: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := strings.TrimSuffix(filepath.Base(frame.File), filepath.Ext(frame.File))
			line := strconv.Itoa(frame.Line)
			return "", fmt.Sprintf("%s:%s", fileName, line)
		},
	})
	log.SetOutput(os.Stdout)
}

func initConfig() {

	viper.AutomaticEnv()
	viper.SetEnvPrefix("PREFIX")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	setupLogging()
}

func init() {
	cobra.OnInitialize(initConfig)
	setParams(rootParams, rootCmd)

}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
