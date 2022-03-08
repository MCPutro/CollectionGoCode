package db_driver

//mysql: github.com/go-sql-driver/mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test?parseTime=true")
	if err != nil {
		fmt.Println("error :", err)
	}

	db.SetMaxIdleConns(10)                  //jumlah koneksi minimum yang dibuat
	db.SetMaxOpenConns(100)                 // jumlah koneksi max yg dibuat
	db.SetConnMaxIdleTime(5 * time.Minute)  // pengaturan berapa lama koneksi yg sudah tidak digunakan akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) //pengaturan brp lama boleh digunakan

	fmt.Println("konek")

	return db
}
