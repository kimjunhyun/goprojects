package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-adodb"
)

func main() {
	mdbFilePath := "./register.mdb"

	db, err := openAccessDb(mdbFilePath)
	if err != nil {
		log.Fatal("open file error:", err)
	}

	tableNames, err := getTableNames(db)
	if err != nil {
		log.Fatal("get table names", err)
	}

	if len(tableNames) <= 0 {
		fmt.Println("no tables found in " + mdbFilePath)
		return
	}
	// print out all the table names
	for index, tableName := range tableNames {
		fmt.Printf("table {%d}: {%s}\n", index+1, tableName)
	}
}

func openAccessDb(mdbFilePath string) (*sql.DB, error) {
	db, err := sql.Open(
		"adodb",
		"Provider=Microsoft.Jet.OLEDB.4.0;Data Source="+mdbFilePath+";")
	return db, err
}

func getTableNames(db *sql.DB) ([]string, error) {
	rows, err := db.Query(
		`SELECT MSysObjects.Name AS table_name
        FROM MSysObjects
        WHERE (((Left([Name],1))<>"~")
        AND ((Left([Name],4))<>"MSys")
        AND ((MSysObjects.Type) In (1,4,6)))
        order by MSysObjects.Name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tableNames := make([]string, 4)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		tableNames = append(tableNames, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tableNames, nil
}
