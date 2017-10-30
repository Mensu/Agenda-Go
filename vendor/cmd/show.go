package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show meetings",
	Long:  `Show meetings between the time`,
	Run: func(cmd *cobra.Command, args []string) {
		startTime, _ := cmd.Flags().GetString("startTime")
		endTime, _ := cmd.Flags().GetString("endTime")
		results, err := service.FindMeetingByTime(startTime, endTime)
		if err == nil {
			for i, meeting := range results {
				fmt.Println("No. %d", i)
				fmt.Println("Title: %s", meeting.Title)
				fmt.Println("Speecher: %s", meeting.Speecher)
				fmt.Println("Participators:")
				for _, participator := range meeting.Participators {
					fmt.Println("             %s", participator)
				}
				fmt.Println("Start Time: %s", meeting.StartTime)
				fmt.Println("End Time: %s", meeting.EndTime)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("startTime", "s", "", "start time for query")
	showCmd.Flags().StringP("endTime", "s", "", "end time for query")
}
