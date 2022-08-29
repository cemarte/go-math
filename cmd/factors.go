/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
  // "reflect"
  "sort"
  "strconv"

	"github.com/spf13/cobra"
)

var knownPrimes = map[int]struct{}{
  2: struct{}{},
  3: struct{}{},
  5: struct{}{},
  7: struct{}{},
  11: struct{}{},
  13: struct{}{},
  17: struct{}{},
  19: struct{}{},
  23: struct{}{},
  29: struct{}{},
}

func isPrime(i int) bool {
  if i<=1 {
    return false
  }

  if _, ok:=knownPrimes[i]; ok {
    return true
  }

  for j:=2; j<=i/2; j++ {
    if i%j==0 {
      return false
    } 
  }

  knownPrimes[i]=struct{}{}
  return true
}



func getPrimeFactors(i int) []int {
  if i==0 || i==1 || i==2 || i==3 {
    return []int{i}
  }

  result:=[]int{}

  for v:=range(knownPrimes) {
    for i%v==0 {
      result = append(result, v)
      i = i/v
    }
    if i==1 {
      break
    }
  }

  // TODO: ordering isn't the same between test runs, check why, I think it relates to array capacity
  sort.Ints(result)
  return result
}
// factorsCmd represents the factors command
var factorsCmd = &cobra.Command{
	Use:   "factors",
	Short: "Given an integer as input, returns its prime factors",
	Long: `Given an integer as input, returns its prime factors`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
      inputs := make([]int, len(args))

      for i := 0; i < len(args); i++ {
        intVal, err := strconv.Atoi(args[i])

        if err != nil {
          panic(err)
        }

        inputs[i] = intVal

        factors:=getPrimeFactors(inputs[i])

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
