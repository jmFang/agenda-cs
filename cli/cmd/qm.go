package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// qmCmd represents the qm command
var qmCmd = &cobra.Command{
	Use:   "qm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("qm called")
		meetings, err := as.ListAllMeetings()
		if err != nil {
			fmt.Println(err)
			log.Infoln(err)
		} else {
			fmt.Println("results:")
			for index, m := range meetings {
				fmt.Printf("%d. %s | %s | %s | %s | %s \n", index+1, m.Title, m.Sponsor, m.Participators, m.Start, m.End)
				//fmt.Println(index+1, m.Title, m.Sponsor, m.Participators, m.Start, m.End)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(qmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// qmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// qmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//qmCmd.Flags().StringP("start", "s", "", "use --start or -s [year-month-day]")
	//qmCmd.Flags().StringP("end", "e", "", "use --end or -e [year-month-day]")
}
