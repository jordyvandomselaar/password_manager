package actions

import (
	"flag"
	"os"
)

// Get available commandline arguments
func Save() (string, string, string, string) {
	name := flag.String("name", "", "Pick a name for this password entry.")
	username := flag.String("username", "", "What is your username?")
	password := flag.String("password", "", "What is your password?")
	description := flag.String("description", "", "Any other information?")

	flag.Parse()

	flagSet := flag.NewFlagSet("save", flag.PanicOnError)

	name := flagSet.String("name", "", "Pick a name for this password entry.")

	flagSet.Parse(os.Args[2:])

	return *name, *username, *password, *description
}
