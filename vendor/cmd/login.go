package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "log in Agenda",
	Long:  `Log in Agenda with username and password`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		fmt.Println("login called by " + username)
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "username")
	loginCmd.Flags().StringP("password", "p", "", "password")
}
