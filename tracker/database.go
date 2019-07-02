package tracker

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3" // Necessary for the "sqlite3" side effect
)

// EntryModel is type to hold data that can be directly added to our database
type EntryModel struct {
	EntryID          string
	Month            int
	Day              int
	Year             int
	Title            string
	TimeModification int
	Note             string
}

var (
	// Database is the active connection to our database
	Database *sql.DB
)

// InitDB creates the file structure necessary for our database
//
// Arguments:
//     None
//
// Returns:
//     (error): returns an error if there is one, otherwise, nil
func InitDB() error {
	dbFileFolder := "CompTimeTracker"

	folderPath, err := checkDbFileFolder(dbFileFolder)
	if err != nil {
		return err
	}

	dbFileName := "compTimeTracker.db"
	fullPath := fmt.Sprintf("%s/%s", folderPath, dbFileName)

	database, err := sql.Open("sqlite3", fullPath)
	if err != nil {
		return err
	}

	s, err := database.Prepare("CREATE TABLE IF NOT EXISTS entries (id INTEGER PRIMARY KEY AUTOINCREMENT, entryID TEXT NOT NULL, month INTEGER NOT NULL, day INTEGER NOT NULL, year INTEGER NOT NULL, title TEXT NOT NULL, timeModification INTEGER NOT NULL, note TEXT)")
	if err != nil {
		return err
	}
	s.Exec()

	Database = database

	return nil
}

// Insert will insert data from EntryModel into our active database
//
// Arguments:
//     None
//
// Augments:
//     *EntryModel
//
// Returns:
//     (error): returns an error if there is one, nil if not
func (E *EntryModel) Insert() error {
	d := Database
	s, err := d.Prepare("INSERT INTO entries (entryID, month, day, year, title, timeModification, note) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	s.Exec(E.EntryID, E.Month, E.Day, E.Year, E.Title, E.TimeModification, E.Note)

	return nil
}

// RemoveEntry will remove an entry from out active database
//
// Arguments:
//     ID (string): the UUID string of the entry to remove
//
// Returns:
//     (error): returns an error if there is on, nil if not
func RemoveEntry(ID string) error {
	d := Database
	q := fmt.Sprintf("DELETE FROM entries WHERE entryID like '%s'", ID)
	s, err := d.Prepare(q)
	if err != nil {
		return err
	}
	s.Exec()

	return nil
}

// GetSingleEntry returns detailed information about a single entry
//
// Arguments:
//     ID (string): the UUID string of the entry to display
//
// Returns:
//     (error): returns an error if there is one, otherwise nil
func GetSingleEntry(ID string) error {
	var (
		entryID          string
		month            int
		day              int
		year             int
		title            string
		timeModification int
		note             string
	)
	d := Database

	q := fmt.Sprintf("SELECT entryID, month, day, year, title, timeModification, note FROM entries WHERE entryID LIKE '%s'", ID)

	rows, err := d.Query(q)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&entryID, &month, &day, &year, &title, &timeModification, &note)
		fmt.Printf(`
ID:    %s
Date:  %d-%d-%d
Title: %s
Time:  %d
Note:  %s
`, entryID, month, day, year, title, timeModification, note)
	}
	return nil
}

// GetAllEntries returns a formatted list of database entries to stout
//
// Arguments:
//      None
//
// Returns:
//     (error): returns an error if there is one, nil if not
func GetAllEntries() error {
	var (
		counter          int
		entryID          string
		month            int
		day              int
		year             int
		title            string
		timeModification int
		note             string
		totalTime        int
	)
	d := Database

	rows, err := d.Query("SELECT entryID, month, day, year, title, timeModification, note FROM entries")
	if err != nil {
		return err
	}
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("|               ENTRY_ID               |    DATE    |                 TITLE                | TIME |                 NOTE                 |")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")
	for rows.Next() {
		rows.Scan(&entryID, &month, &day, &year, &title, &timeModification, &note)

		if len(title) > 35 {
			title = title[:33] + "..."
		} else if len(title) < 36 {
			for len(title) < 36 {
				title = title + " "
			}
		}

		if len(note) > 35 {
			note = note[:33] + "..."
		} else if len(note) == 0 {
			for len(note) < 36 {
				note = note + " "
			}
		} else if len(note) < 36 {
			for len(note) < 36 {
				note = note + " "
			}
		}

		fmt.Printf("| %s | %02d-%02d-%d | %s |  %d  | %s |\n", entryID, month, day, year, title, timeModification, note)
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")
		counter = counter + 1
		totalTime = totalTime + timeModification
	}
	fmt.Printf("TOTAL:\nEntries: %d | Comp Time (in Minutes): %v | Comp Time (in Hours): %v | Comp Time (in Days): %v |\n", counter, totalTime, float32(totalTime)/float32(60), float32(totalTime)/float32(1440))
	return nil
}

// GetTotal returns an int of total mintues of all entries (+ and -)
//
// Arguments:
//     None
//
// Returns:
//     (int): The total
//     (error): returns an error if there is one, nil if not
func GetTotal() (int, error) {
	var (
		timeModification int
		total            int
	)

	d := Database
	rows, err := d.Query("SELECT timeModification FROM entries")
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		rows.Scan(&timeModification)
		total = total + timeModification
	}
	return total, nil
}

func checkDbFileFolder(folderName string) (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := fmt.Sprintf("%s/%s", userHome, folderName)

	stat, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(fullPath, 0775)
			if err != nil {
				return "", err
			}
			return fullPath, nil
		}
		return "", err
	}
	if stat.IsDir() {
		return fullPath, nil
	}
	return "", errors.New("Schrodinger's Directory")
}
