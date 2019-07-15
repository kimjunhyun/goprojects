package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
	_ "github.com/mattn/go-adodb"

	"github.com/gonutz/w32"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

var provider string

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
}

func main() {
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, w32.SW_MINIMIZE)
		}
	}

	for {
		go GetDB()
		time.Sleep(time.Second)
	}
}

func GetDB() {

	var ID int
	var 호실 sql.NullString
	var MOD1 sql.NullString
	var MOD2 sql.NullString
	var MOD3 sql.NullString
	var 층 sql.NullString
	var 객실종류 sql.NullString
	var 대실종류 sql.NullString
	var 객실상태 sql.NullString
	var 대실일 sql.NullString
	var 대실시각 sql.NullString
	var 종료일 sql.NullString
	var 종료시각 sql.NullString
	var 차량번호 sql.NullString
	var 문상태 bool
	var 대실키 sql.NullString
	var 청소키 bool
	var AutoLock bool
	var 판매기 bool
	var lock_state sql.NullString
	var 대실요금 int
	var 현금체크1 sql.NullString
	var 숙박요금 int
	var 현금체크2 sql.NullString
	var 추가요금 int
	var 현금체크3 sql.NullString
	var 입금종류 sql.NullString
	var ImgTop sql.NullString
	var ImgLeft sql.NullString
	var ImgWidth sql.NullString
	var ImgHeight sql.NullString
	var 판매형태 sql.NullString

	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// if _, err := os.Stat(cfg.Section("DB").Key("source").String()); err != nil {
	// 	fmt.Println("put here source database named '" + cfg.Section("DB").Key("source").String() + "'.")
	// 	return
	// }
	// if _, err := os.Stat(cfg.Section("DB").Key("target").String()); err != nil {
	// 	fmt.Println("put here target database named '" + cfg.Section("DB").Key("target").String() + "'.")
	// 	return
	// }

	var bufs bytes.Buffer
	wr := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())
	wr.Write([]byte(cfg.Section("DB").Key("target").String()))
	wr.Close()
	//convVal := bufs.String()
	fmt.Println("Provider=Microsoft.ACE.OLEDB.12.0;Data Source=" + cfg.Section("DB").Key("target").String() + ";")
	return

	db, err := sql.Open("adodb", "Provider=Microsoft.ACE.OLEDB.12.0;Data Source="+cfg.Section("DB").Key("source").String()+";")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	dbtarget, err := sql.Open("adodb", "Provider=Microsoft.ACE.OLEDB.12.0;Data Source="+cfg.Section("DB").Key("target").String()+";")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbtarget.Close()

	rows, err := db.Query("SELECT * FROM ROOM_INF")

	// Prepared Statement 생성
	stmt, err := dbtarget.Prepare("UPDATE ROOM_INF SET 객실상태=? WHERE 호실=?")
	checkError(err)
	defer stmt.Close()

	for rows.Next() {
		err := rows.Scan(&ID, &호실, &MOD1, &MOD2, &MOD3, &층, &객실종류, &대실종류, &객실상태, &대실일, &대실시각, &종료일, &종료시각, &차량번호, &문상태, &대실키, &청소키, &AutoLock, &판매기, &lock_state, &대실요금, &현금체크1, &숙박요금, &현금체크2, &추가요금, &현금체크3, &입금종류, &ImgTop, &ImgLeft, &ImgWidth, &ImgHeight, &판매형태)
		if err != nil {
			log.Fatal(err)
		}

		str := 객실상태.String
		str1 := 호실.String
		fmt.Printf("호실 '%s', 객실상태: '%s'\n", str1, str)

		if str == "사용중지" {
			str = ""
		}
		// _, err = stmt.Exec(str, 호실.String) //Placeholder 파라미터 순서대로 전달
		strq := fmt.Sprintf("UPDATE ROOM_INF SET 객실상태 = '%s', 대실일='%s', 대실시각='%s', 대실요금=%d, 숙박요금=%d, 입금종류='%s' WHERE 호실='%s'", str, 대실일.String, 대실시각.String, 대실요금, 숙박요금, 입금종류.String, str1)
		// fmt.Println(strq)
		_, err = dbtarget.Exec(strq)

		checkError(err)

		//fmt.Println(roomnum, 객실상태)
	}

	// sqls := []string{
	// 	//"select * from languages",
	// 	"DROP TABLE languages",
	// 	"CREATE TABLE languages (id text not null primary key, name text)",
	// }
	// for _, sql := range sqls {
	// 	_, err = db.Exec(sql)
	// 	if err != nil {
	// 		fmt.Printf("!!! %q: %s\n", err, sql)
	// 		return
	// 	}
	// }
	// fmt.Printf("!!! &&&")
	// // tx, err := db.Begin()
	// // if err != nil {
	// // 	fmt.Println(err)
	// // 	return
	// // }

	// // stmt, err := tx.Prepare("insert into languages(id, name) values("append", "ss"")
	// sss := "insert into languages(id, name) values('append', 'ss')"
	// _, err = db.Exec(sss)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //defer stmt.Close()

	// fmt.Printf("!!! &&&1")

	// _, err = stmt.Exec("en", "English")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("!!! &&&2")
	// _, err = stmt.Exec("fr", "French")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// _, err = stmt.Exec("de", "German")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// _, err = stmt.Exec("es", "Spanish")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("!!! &&&3")

	// tx.Commit()
	fmt.Printf("!!! &&&4")
}
