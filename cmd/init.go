/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/GolangSriLanka/glc/cmd/prompt"
	"github.com/GolangSriLanka/glc/internal"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var (
	typeFlag string
	initCmd  = &cobra.Command{
		Use:   "init",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			internal.CreateConfigDir()
			initGLC()
		},
	}
)

func init() {
	initCmd.Flags().StringVarP(&typeFlag, "type", "t", "", "--type=<TYPE> or -t <TYPE>")
	rootCmd.AddCommand(initCmd)
}

func initGLC()  {
	if typeFlag == "" {
		fmt.Println("No flag provided")
		prompt.SelectProviderWithPrompt()
		return
	}

	fmt.Println(typeFlag)
}
