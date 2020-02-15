package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCMDInput = struct {
	cfPath       string
	mode         string
	checkImpTree bool
	rootPath     string
	skipFol      string
	refDataPath  string
	schemaPath   string
	isExport     bool
}{}

var initCMD = &cobra.Command{
	Use:   "init",
	Short: "init the storage server of webserver",
	Long:  `webserver server start run `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webserver Static Site Generator v0.9 -- HEAD")
		fmt.Println("write out config")

	},
}

func init() {
	callPath, _ := os.Getwd()
	initCMD.Flags().StringVarP(
		&initCMDInput.cfPath,
		"conf", "c",
		filepath.Join(callPath, "config.toml"),
		"start server with specific config file")

	initCMD.Flags().BoolVarP(
		&initCMDInput.isExport,
		"exportSchema", "x",
		false,
		"export the createCollecction validator.$jsonSchema build json")
	initCMD.Flags().StringVarP(
		&initCMDInput.schemaPath,
		"schema", "m",
		filepath.Join(callPath, "doc", "log", "create_schema.json"),
		"start server with specific config file")
	rootCmd.AddCommand(initCMD)
}
