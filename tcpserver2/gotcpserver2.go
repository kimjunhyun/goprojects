package main

import (
	"database/sql"
	"fmt"
	"net"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "127.0.0.1"
	database = "test"
	user     = "root"
	password = "admin"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func requestHandler(c net.Conn) {
	data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성

	for {
		n, err := c.Read(data) // 클라이언트에서 받은 데이터를 읽음
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) // 데이터 출력

		_, err = c.Write(data[:n]) // 클라이언트로 데이터를 보냄
		if err != nil {
			fmt.Println(err)
			return
		}

		db(string(data[:n])) //데이타를 저장함

	}
}

func main() {
	ln, err := net.Listen("tcp", ":8000") // TCP 프로토콜에 8000 포트로 연결을 받음
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close() // main 함수가 끝나기 직전에 연결 대기를 닫음

	for {
		conn, err := ln.Accept() // 클라이언트가 연결되면 TCP 연결을 리턴
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

		go requestHandler(conn) // 패킷을 처리할 함수를 고루틴으로 실행
	}
}

func db(msg string) {

	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")

	// Drop previous table of same name if one exists.
	_, err = db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed).")

	// Create table.
	_, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table.")

	// Insert some data into table.
	sqlStatement, err := db.Prepare("INSERT INTO inventory (name, quantity) VALUES (?, ?);")
	res, err := sqlStatement.Exec("banana", 150)
	checkError(err)
	rowCount, err := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	res, err = sqlStatement.Exec("orange", 154)
	checkError(err)
	rowCount, err = res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	if len(msg) < 5 {
		res, err = sqlStatement.Exec("apple", msg)
		checkError(err)
		rowCount, err = res.RowsAffected()
		fmt.Printf("Inserted %d row(s) of data.\n", rowCount)
		fmt.Println("Done.")
	}
}
