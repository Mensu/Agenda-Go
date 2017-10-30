package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var quitMeetingCmd = &cobra.Command{
	Use:   "quit-meeting",
	Short: "quit a meeting",
	Long:  `quit a meeting by participator`,
	Run: func(cmd *cobra.Command, args []string) {
		title := cmd.Flags().GetString("title")
		err := service.DeleteFromMeeting(title)
		if err == nil {
			fmt.Println("Quited the meeting %s", title)
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(quitMeetingCmd)
	quitMeetingCmd.Flags().StringP("title", "t", "", "the title of the meeting")
}
