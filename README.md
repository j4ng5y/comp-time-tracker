# Comp Time Tracker

This is a small little app that just helps me track my comp time while working on contracts.

## Installation

### Linux

`sudo curl -o /usr/local/bin/ctt https://github.com/j4ng5y/comp-time-tracker/releases/download/2019.07.01/ctt-2019.07.01-linux-amd64`

### OSX

`sudo curl -o /usr/local/bin/ctt https://github.com/j4ng5y/comp-time-tracker/releases/download/2019.07.01/ctt-2019.07.01-darwin-amd64`

### Windows

See building from source

## Building From Source

1. Have Go installed (you can get it [here](https://golang.org/dl/))
2. Clone This Repository (`git clone https://github.com/j4ng5y/comp-time-tracker $GOROOT/src/github.com/j4ng5y/comp-time-tracker`)
3. Navigate into the repository (`cd $GOROOT/src/github.com/j4ng5y/comp-time-tracker`)
4. Build the tool (`CGO_ENABLE=1 go build -a -installsuffix cgo -o bin/ctt`)
5. Install the tool:
    * `ln -s $GOROOT/src/github.com/j4ng5y/comp-time-tracker/bin/ctt /usr/local/bin/ctt`
    or
    * `go install`

## Usage

```text
A small app to track comp time

Usage:
  cct [flags]
  cct [command]

Available Commands:
  delete      delete a comp time entry
  help        Help about any command
  new         create a new comp time entry
  view        view comp time entries

Flags:
  -h, --help      help for cct
      --version   version for cct

Use "cct [command] --help" for more information about a command.
```

## Examples

### Adding a new time entry

#### new Usage

```text
create a new comp time entry

Usage:
  cct new [flags]

Flags:
  -d, --day int        the day of the entry
  -h, --help           help for new
  -m, --month int      the month of the entry
  -n, --note string    a note for the entry
  -T, --time int       the amount of time to add (positive integer) or subtract (negative integer) from the running total
  -t, --title string   the title of the entry to create
  -y, --year int       the year of the entry
```

#### new Commands Examples

`ctt new -t "<title of entry>" -T <minutes to add> -m <month of entry> -d <day of entry> -y <year of entry> -n "<optional: note>"`

To add a new 30 minute entry on July 1, 2019, it would like this:

`ctt new -t "This is a test entry" -T 30 -m 07 -d 01 -y 2019 -n "This test entry is used just to test the functionality of the ctt tool."`

You will see something to the effect of this:

`2019/07/02 11:26:01 Entry {955c8c75-0d7d-48f2-a534-866abcdd1f31 7 1 2019 This is a test entry 30 This test entry is used just to test the functionality of the ctt tool.} added`

### Viewing entries

#### view Usage

```text
view comp time entries

Usage:
  cct view [flags]

Flags:
  -h, --help                  help for view
  -s, --single-entry string   view a single entry
  -D, --total-days            use this flag to view running time in days
  -H, --total-hours           use this flag to view running time in hours
  -M, --total-minutes         use this flag to view running time in minutes
  -t, --total-only            use this flag if you only want to output the running total
```

#### view Command Examples

*View All Entries*
`ctt view`

This will output something like this (expanding on the previous example):

```text
------------------------------------------------------------------------------------------------------------------------------------------
|               ENTRY_ID               |    DATE    |                 TITLE                | TIME |                 NOTE                 |
------------------------------------------------------------------------------------------------------------------------------------------
| 955c8c75-0d7d-48f2-a534-866abcdd1f31 | 07-01-2019 | This is a test entry                 |  30  | This test entry is used just to t... |
------------------------------------------------------------------------------------------------------------------------------------------
TOTAL:
Entries: 1 | Comp Time (in Minutes): 30 | Comp Time (in Hours): 0.5 | Comp Time (in Days): 0.020833334 |
```

*View One Entry*
`ctt view -s "<UUID>"`

This will output something like this (expanding on the previous example):

```text
ID:    955c8c75-0d7d-48f2-a534-866abcdd1f31
Date:  7-1-2019
Title: This is a test entry
Time:  30
Note:  This test entry is used just to test the functionality of the ctt tool.
```

*View Totals Only*
**View In Minutes**: `ctt view -t -M`

Which outputs something like this (expanding on the previous example):

```text
2019/07/02 11:35:30 The current running total of all comp time is: 30 minutes
```

**View In Hours**: `ctt view -t -H`

Which outputs something like this (expanding on the previous example):

```text
2019/07/02 11:35:30 The current running total of all comp time is: 0 hours
```

**View In Days**: `ctt view -t -D`

Which outputs something like this (expanding on the previous example):

```text
2019/07/02 11:36:21 The current running total of all comp time is: 0 days
```

### Delete Entries

#### delete Usage

```text
delete a comp time entry

Usage:
  cct delete [flags]

Flags:
  -h, --help        help for delete
  -i, --id string   the ID of the entry to delete
```

#### delete Command Examples

`ctt delete -i "<UUID>"`

Which outputs something like this (expanding on the previous example):

```text
2019/07/02 11:39:01 Successfully removed entry with ID: 955c8c75-0d7d-48f2-a534-866abcdd1f31
```
