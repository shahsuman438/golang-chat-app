/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"client/controller"
	"client/models"
	"github.com/spf13/cobra"
	"log"
)

type register struct {
	Email    string
	UserName string
	Password string
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a user",
	Long:  `Resiter a user cmd email username password`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			log.Printf("Error:- use cmd ./Client register <email> <user name> <password>\n")
		} else {
			// values := map[string]string{"email": args[0], "userName": args[1], "password": args[2]}
			var values models.Register
			values.Email = args[0]
			values.UserName = args[1]
			values.Password = args[2]
			controller.Register(values)

		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
