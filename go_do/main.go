/*
Copyright © 2022 Chris Bailey

*/
package main

import (
	"github.com/xchrisbailey/go_do/cmd"
	"github.com/xchrisbailey/go_do/data"
)

func main() {
	db, _ := data.DbConnect()
	data.InitDatabase(db)
	defer db.Close()

	cmd.Execute()

}
