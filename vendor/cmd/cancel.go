package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "cancel the current account of Agenda",
	Long:  `cancel the current account of Agenda`,
	Run: func(cmd *cobra.Command, args []string) {
		err := service.DeleteUser()
		if err == nil {
			fmt.Println("Cancelled successfully")
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(cancelCmd)
}
