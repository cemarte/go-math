/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// isPrimeCmd represents the isPrime command
var isPrimeCmd = &cobra.Command{
	Use:   "isPrime",
	Short: "Given an interger, returns if the number is prime",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			panic("missing args")
		}

		inputs := make([]int, len(args))
		for i := 0; i < len(args); i++ {
			intVal, err := strconv.Atoi(args[i])

			if err != nil {
				panic(err)
			}

			inputs[i] = intVal

			result := isPrime(inputs[i])

			fmt.Printf("%d is Prime? %t\n", intVal, result)

		}
	},
}

func init() {
	rootCmd.AddCommand(isPrimeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// isPrimeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// isPrimeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
