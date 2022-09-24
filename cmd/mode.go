/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/cemarte/gre/go-math/lib"
	"github.com/spf13/cobra"
)

// modeCmd represents the mode command
var modeCmd = &cobra.Command{
	Use:   "mode",
	Short: "Given a list of numbers, returns the mode",
	Long:  `Return the term that appears most frequently in a set`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			panic("missing arguments")
		}

		inputs := make([]int, len(args))

		for i := 0; i < len(args); i++ {
			intVal, err := strconv.Atoi(args[i])

			if err != nil {
				panic(err)
			}

			inputs[i] = intVal

		}
		result, hasMode := lib.GetMode(inputs)

		fmt.Printf("Has Mode? %t \t Mode: %v \n", hasMode, result)
	},
}

func init() {
	rootCmd.AddCommand(modeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
