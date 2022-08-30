/*Package cmd:
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// factorsCmd represents the factors command
var factorsCmd = &cobra.Command{
	Use:   "factors",
	Short: "Given an integer as input, returns its prime factors",
	Long:  `Given an integer as input, returns its prime factors`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			inputs := make([]int, len(args))

			for i := 0; i < len(args); i++ {
				intVal, err := strconv.Atoi(args[i])

				if err != nil {
					panic(err)
				}

				inputs[i] = intVal

				factors := getPrimeFactors(inputs[i])

				fmt.Println(factors)

			}
		}
	},
}

func init() {
	rootCmd.AddCommand(factorsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// factorsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// factorsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
