package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	version = "2019.06.29"

	compTimeTrackerCmd = &cobra.Command{
		Use:     "cct",
		Version: version,
		Short:   "A small app to track comp time",
		Long:    "",
		Run:     func(ccmd *cobra.Command, args []string) {},
	}
)

// Execute executes the CLI
//
// Arguments:
//     None
//
// Returns:
//     None
func Execute() {
	err := compTimeTrackerCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
