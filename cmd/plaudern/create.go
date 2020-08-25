package main

import (
	"encoding/json"
	"github.com/bst27/plaudern/internal/configuration"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var (
	createCmd = &cobra.Command{
		Use:   "create::defaultConfig",
		Short: "Write default config to a file.",
		Long:  "Write the default config to a file in JSON format.",
		Run: func(cmd *cobra.Command, args []string) {
			json, err := json.MarshalIndent(configuration.GetDefault(), "", "  ")
			if err != nil {
				log.Fatalln(err)
			}

			err = ioutil.WriteFile(cmd.Flag("path").Value.String(), json, 0644)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
)

func init() {
	createCmd.Flags().String("path", "plaudern-config.json", "Target file to save default config.")
	err := createCmd.MarkFlagRequired("path")
	if err != nil {
		log.Fatalln(err)
	}

	rootCmd.AddCommand(createCmd)
}
