package cmd

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/j4ng5y/comp-time-tracker/tracker"
	"github.com/spf13/cobra"
)

var (
	entryID            string
	entryTitle         string
	entryMonth         int
	entryDay           int
	entryYear          int
	entryTimeInMinutes int
	entryNote          string
	totalOnly          bool
	totalMinutes       bool
	totalHours         bool
	totalDays          bool

	newEntryCmd = &cobra.Command{
		Use:   "new",
		Short: "create a new comp time entry",
		Long:  "",
		Run:   newCompTimeEntry,
	}

	deleteEntryCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a comp time entry",
		Long:  "",
		Run:   deleteCompTimeEntry,
	}

	viewEntriesCmd = &cobra.Command{
		Use:   "view",
		Short: "view comp time entries",
		Long:  "",
		Run:   viewCompTimeEntries,
	}
)

func init() {
	nowM, err := strconv.Atoi(time.Now().Month().String())
	if err != nil {
		nowM = 00
	}
	nowD, err := strconv.Atoi(string(time.Now().Day()))
	if err != nil {
		nowD = 00
	}
	nowY, err := strconv.Atoi(string(time.Now().Year()))
	if err != nil {
		nowY = 0000
	}

	compTimeTrackerCmd.AddCommand(newEntryCmd)
	compTimeTrackerCmd.AddCommand(deleteEntryCmd)
	compTimeTrackerCmd.AddCommand(viewEntriesCmd)

	newEntryCmd.Flags().StringVarP(&entryTitle, "title", "t", "", "the title of the entry to create")
	newEntryCmd.Flags().IntVarP(&entryTimeInMinutes, "time", "T", 0, "the amount of time to add (positive integer) or subtract (negative integer) from the running total")
	newEntryCmd.Flags().IntVarP(&entryMonth, "month", "m", nowM, "the month of the entry")
	newEntryCmd.Flags().IntVarP(&entryDay, "day", "d", nowD, "the day of the entry")
	newEntryCmd.Flags().IntVarP(&entryYear, "year", "y", nowY, "the year of the entry")
	newEntryCmd.Flags().StringVarP(&entryNote, "note", "n", "", "a note for the entry")
	newEntryCmd.MarkFlagRequired("title")
	newEntryCmd.MarkFlagRequired("time")
	newEntryCmd.MarkFlagRequired("month")
	newEntryCmd.MarkFlagRequired("day")
	newEntryCmd.MarkFlagRequired("year")

	deleteEntryCmd.Flags().StringVarP(&entryID, "id", "i", "", "the ID of the entry to delete")
	deleteEntryCmd.MarkFlagRequired("id")

	viewEntriesCmd.Flags().StringVarP(&entryID, "single-entry", "s", "", "view a single entry")
	viewEntriesCmd.Flags().BoolVarP(&totalOnly, "total-only", "t", false, "use this flag if you only want to output the running total")
	viewEntriesCmd.Flags().BoolVarP(&totalMinutes, "total-minutes", "M", false, "use this flag to view running time in minutes")
	viewEntriesCmd.Flags().BoolVarP(&totalHours, "total-hours", "H", false, "use this flag to view running time in hours")
	viewEntriesCmd.Flags().BoolVarP(&totalDays, "total-days", "D", false, "use this flag to view running time in days")
}

func newCompTimeEntry(ccmd *cobra.Command, args []string) {
	var E tracker.EntryModel
	E.EntryID = uuid.New().String()
	E.Month = entryMonth
	E.Day = entryDay
	if entryYear <= 1989 {
		log.Fatalf("the year %d is out of bounds; verify that your year is four digits long\n", entryYear)
		os.Exit(1)
	}
	if entryYear >= 2050 {
		log.Fatalf("the year %d is out of bounds; verify that your year is four digits long\n", entryYear)
		os.Exit(1)
	}
	E.Year = entryYear
	E.Title = entryTitle
	E.TimeModification = entryTimeInMinutes
	E.Note = entryNote
	err := E.Insert()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Entry %v added", E)
}

func deleteCompTimeEntry(ccmd *cobra.Command, args []string) {
	err := tracker.RemoveEntry(entryID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully removed entry with ID: %s\n", entryID)
}

func viewCompTimeEntries(ccmd *cobra.Command, args []string) {
	if entryID != "" {
		err := tracker.GetSingleEntry(entryID)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if totalOnly {
		t, err := tracker.GetTotal()
		if err != nil {
			log.Fatal(err)
		}
		if totalMinutes {
			log.Printf("The current running total of all comp time is: %v minutes\n", t)
			return
		} else if totalHours {
			log.Printf("The current running total of all comp time is: %v hours\n", t/60)
			return
		} else if totalDays {
			log.Printf("The current running total of all comp time is: %v days\n", t/1440)
			return
		} else {
			log.Println("one of '--total-minutes (-M)', '--total-hours (-H)', or '--total-days (-D)' is required")
			return
		}
	}
	err := tracker.GetAllEntries()
	if err != nil {
		log.Fatal(err)
	}
}
