package main

import (
	"fmt"
	db "github.com/MCPutro/CollectionGoCode/modul1/init-package/database"
)

func main() {
	fmt.Println(">", db.GetConnection())
}
