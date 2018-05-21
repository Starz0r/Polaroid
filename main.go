package main

import (
	"polaroid/cmd"
	
	"os"
)

func main() {
	if os.Args[1] == "install" {
		cmd.DatabaseMigrations()
	}
}