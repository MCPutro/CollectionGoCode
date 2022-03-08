package test

import (
	"context"
	"database/sql"
	"fmt"
	db_driver "github.com/MCPutro/CollectionGoCode/modul8-database/db-driver"
	"testing"
	"time"
)

func TestInsertSederhana(t *testing.T) {
	db := db_driver.GetConnection()

	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx,
		"insert into customer (id, name) values ('pertama', 'kesatu');")

	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Println("berhasi")

}

func TestSelect(t *testing.T) {
	db := db_driver.GetConnection()

	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx,
		"select id, name, email, balance, rating, create_date, birth_date, married from customer")

	if err != nil {
		fmt.Println("error select :", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name, email sql.NullString
		var balance int32
		var rating float32
		var createDate time.Time
		var birthDate sql.NullTime
		var marr bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &createDate, &birthDate, &marr) //susunan attribute harus sama dengan query SQL nya
		if err != nil {
			fmt.Println("error print :", err)
		}

		fmt.Println("id :", id)
		fmt.Println("name :", name)
		if email.Valid {
			fmt.Println("email :", email.String)
		}
		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		fmt.Println("createDate :", createDate)
		fmt.Println("birthDate :", birthDate)
		fmt.Println("married :", marr)
		fmt.Println("----------------------------")
	}
}

func TestInsertParam(t *testing.T) {
	db := db_driver.GetConnection()
	defer db.Close()

	id := "email@gmail.com"
	name := "lupa cuy"

	script := "insert into comments ( email, comment) values ( ?, ?);"

	result, err := db.ExecContext(context.Background(), script, id, name)

	if err != nil {
		fmt.Println("error1 :", err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("error2 :", err)
	}
	fmt.Println("berhasil insert dengan id :", insertId)

}

//jika ingin melakukan query yg sama berulang kali lebih baik dengan prepare statement
func TestPrepareStatement(t *testing.T) {
	db := db_driver.GetConnection()
	defer db.Close()

	ctx := context.Background()

	sql_query := "insert into comments ( email, comment) values ( ?, ?);"

	stmt, err := db.PrepareContext(ctx, sql_query)
	if err != nil {
		panic(err)
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	for i := 0; i < 10; i++ {
		email := fmt.Sprint("test", i, "@gmail.com")
		comment := fmt.Sprint("comment ke-", i)

		queryContext, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := queryContext.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("id :", id)

	}
}

//bisa set kapan harus rollback
func TestTransaction(t *testing.T) {
	db := db_driver.GetConnection()
	defer db.Close()
	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	//start transaksi
	sql_query := "insert into comments ( email, comment) values ( ?, ?);"
	for i := 0; i < 10; i++ {
		email := fmt.Sprint("test", i, "@gmail.com")
		comment := fmt.Sprint("comment ke-", i)

		execContext, err := tx.ExecContext(ctx, sql_query, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := execContext.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("id :", id)

	}
	//end transaksi

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		panic(err)
	}
}
