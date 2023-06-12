package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	server := "10.199.14.47"
	port := 1433
	user := "kemahasiswaan_app"
	password := "kemahasiswaan_app123#"
	database := "KEMAHASISWAAN"

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)

	db, err := sql.Open("sqlserver", connString)
	defer db.Close()
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	tsql := fmt.Sprintf("SELECT CONVERT(nvarchar(36), id_mhs), nama FROM magang.registrasi_magang ORDER BY id_registrasi_magang OFFSET 0 ROW FETCH NEXT 10 ROWS ONLY")
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		var name, id string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("ID: %v, name: %s\n", ("s" + id), name)
		count++
	}

	fmt.Println(count)

}
