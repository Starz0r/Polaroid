package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func DatabaseMigrations() {
	reader := bufio.NewReader(os.Stdin)

	// Select Database Type
	//* Only supports the same database types as the ORM
	//TODO: Handle outside cases that aren't supported types
	fmt.Print(">>>Polaroid - Choose a database type: ")
	dbType, _ := reader.ReadString('\n')

	// Select Database Table Name
	fmt.Print(">>>Polaroid - What should the table be named?: ")
	dbName, _ := reader.ReadString('\n')

	// Select Database User
	fmt.Print(">>>Polaroid - What is the database username: ")
	dbUser, _ := reader.ReadString('\n')

	// Select Database User Password
	fmt.Print(">>>Polaroid - Password to that account?: ")
	dbPass, _ := reader.ReadString('\n')
}
