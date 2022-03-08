package main

import (
	db "Dasar2Go/modul1/init-package/database"
	"fmt"
)

func main() {
	fmt.Println(">", db.GetConnection())
}
