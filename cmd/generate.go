/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// type cmdFlags struct {
// 	enterprise   string
// 	organization string
// 	repository   string
// }

var enterprise string
var organization string
var repository string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		if len(args) == 0 {
			runInteractiveMode()
		}

		// var err error
		// var client api.GQLClient

		// client, err = gh.GQLClient(&api.ClientOptions{
		// 	Headers: map[string]string{
		// 		"Accept": "application/vnd.github.hawkgirl-preview+json",
		// 	},
		// })

		// if err != nil {
		// 	return err
		// }

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().StringVarP(&enterprise, "enterprise", "e", "", "Generate GHAS CSV for an enterprise.")
	generateCmd.Flags().StringVarP(&organization, "organization", "o", "", "Generate GHAS CSV for an organization.")
	generateCmd.Flags().StringVarP(&repository, "repository", "r", "", "Generate GHAS CSV for an repository.")
	// generateCmd.MarkFlagsRequiredTogether("repository", "organization") // To Do: Repo requires org, but org doesn't require repo
}

// A function that prints the word hello
func runInteractiveMode() {
	fmt.Println("Running in interactive mode")
}

// func runEmpty() {
// 	fmt.Println('Running Empty')
// }

// A function that checks is a variable is an empty array {
// func (a []string) IsEmpty() bool {
// 	return len(a) == 0
// }
