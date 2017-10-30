package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear all meeting by speecher",
	Long:  `clear all meeting by speecher`,
	Run: func(cmd *cobra.Command, args []string) {
		title := cmd.Flags().GetString("title")
		err := service.DeleteAllMeeting(title)
		if err == nil {
			fmt.Println("Cleared all the meetings")
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(clearCmd)
}
