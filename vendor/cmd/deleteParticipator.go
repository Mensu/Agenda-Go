package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var deleteParticipatorCmd = &cobra.Command{
	Use:   "delete-participator",
	Short: "delete a participator from a meeting",
	Long:  `delete a existed participator from a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		title := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("participator")
		err := service.DeleteParticipatorToMeeting(title, participators)
		if err == nil {
			fmt.Println("Deleted participator from the meeting %s", title)
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.addCommand(deleteParticipatorCmd)
	deleteParticipatorCmd.Flags().StringP("title", "t", "", "the title of the meeting")
	deleteParticipatorCmd.Flags().StringArrayP("participator", "p", "", "the participator of the meeting")
}
