/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	jwtconverter "github.com/willjrcom/jwt-golang/internal"
)

// createTokenCmd represents the createToken command
var createTokenCmd = &cobra.Command{
	Use:   "create-token",
	Short: "Create your jwt token",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create-token called")
		token, err := jwtconverter.CreateToken()

		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(token)
	},
}

func init() {
	rootCmd.AddCommand(createTokenCmd)
}
