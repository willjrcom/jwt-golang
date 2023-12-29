package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	jwtconverter "github.com/willjrcom/jwt-golang/internal"
)

var readTokenCmd = &cobra.Command{
	Use:   "read-token",
	Short: "Read jwt token",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Read JWT token")
		token, err := cmd.Flags().GetString("token")

		if err != nil {
			log.Fatal(err)
			return
		}

		if token == "" {
			log.Fatalln("token not found")
			return
		}

		data, _ := jwtconverter.ReadToken(token)

		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(data)
	},
}

func init() {
	readTokenCmd.Flags().StringP("token", "t", "", "token for convertion")
	rootCmd.AddCommand(readTokenCmd)
}
