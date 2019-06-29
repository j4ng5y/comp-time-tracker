package cmd

import (
	"log"
	"github.com/spf13/cobra"
)

var (
	version = "2019.06.29"

	compTimeTrackerCmd = &cobra.Command{
		Use: "comp-time-tracker",
		Version: version,
		Short: "A small app to track comp time",
		Long: "",
		Run: func(ccmd *cobra.Command, args []string){},
	}
)

func Execute() {
	err := compTimeTrackerCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}