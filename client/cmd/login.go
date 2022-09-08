/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"client/controller"
	"client/interfaces"
	"log"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login user with cmd email password",
	Long:  `login for user auth`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Printf("Error:- use cmd ./Client login <email> <password>\n")
		} else {
			var loginData interfaces.Login
			loginData.Email = args[0]
			loginData.Password = args[1]
			controller.Login(loginData)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
