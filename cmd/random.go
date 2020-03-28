package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generate some random strings",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		alphaLowerCaseChars := "abcdefghijklmnopqrstuvwxyz"
		alphaUpperCaseChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numericChars := "0123456789"
		specialChars := "!\"#$%&'()*+,-./:;<=>?@[]^_{|}~"
		alpha, _ := cmd.Flags().GetBool("alpha")
		uppercase, _ := cmd.Flags().GetBool("uppercase")
		lowercase, _ := cmd.Flags().GetBool("lowercase")
		special, _ := cmd.Flags().GetBool("special")
		digits, _ := cmd.Flags().GetBool("digits")
		noHeaders, _ := cmd.Flags().GetBool("no-headers")
		finalStr := ""
		if !alpha && !uppercase && !lowercase && !digits && !special {
			finalStr += alphaLowerCaseChars + alphaUpperCaseChars + specialChars + numericChars
		} else {
			if alpha && !uppercase && !lowercase {
				finalStr += alphaUpperCaseChars + alphaLowerCaseChars
			} else if uppercase {
				finalStr += alphaUpperCaseChars
			} else if lowercase {
				finalStr += alphaLowerCaseChars
			}
			if special {
				finalStr += specialChars
			}
			if digits {
				finalStr += numericChars
			}
		}
		number, _ := cmd.Flags().GetInt("number")
		length, _ := cmd.Flags().GetInt("length")
		if len(args) > 0 {
			number = len(args)
		}
		rand.Seed(time.Now().Unix())
		for i := 0; i < number; i++ {
			randStr := ""
			for j := 0; j < length; j++ {
				randStr += string(finalStr[rand.Intn(len(finalStr))])
			}
			if noHeaders {
				if i > 0 {
					fmt.Print("\n")
				}
				fmt.Printf("%v", randStr)
			} else if len(args) > 0 {
				fmt.Printf("%v: %v\n", args[i], randStr)
			} else {
				fmt.Printf("%3d: %v\n", i+1, randStr)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.Flags().Bool("no-headers", false, "No headers")
	randomCmd.Flags().IntP("number", "n", 1, "Number of strings to generate")
	randomCmd.Flags().Bool("alpha", false, "Use alpha characters")
	randomCmd.Flags().Bool("uppercase", false, "Use uppercases")
	randomCmd.Flags().Bool("lowercase", false, "Use upperlower cases")
	randomCmd.Flags().Bool("digits", false, "Use digits")
	randomCmd.Flags().Bool("special", false, "Use special")
	randomCmd.Flags().IntP("length", "l", 10, "Length of each string")
}
