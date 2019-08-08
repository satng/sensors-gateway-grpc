package taosTool

import (
	"database/sql"
	"fmt"
	_ "github.com/satng/sensors-gateway-grpc/taosSql"
	"log"
	"time"
)

const (
	CONNDB  = "%s:%s@/tcp(%s)/%s"
	DRIVER  = "taosSql"
	DBhost  = "127.0.0.1"
	DBuser  = "root"
	DBpass  = "taosdata"
	DBname  = "sensors_test"
	STable1 = "sensors_data"
	STable2 = "gps_data"
)

func Test() bool {
	// open connect to taos server
	connStr := fmt.Sprintf(CONNDB, DBuser, DBpass, DBhost, DBname)
	db, err := sql.Open(DRIVER, connStr)
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
		return false
	} else {
		fmt.Printf("Open database ok\n")
		return true
	}
	defer db.Close()
}

func insert_data(db *sql.DB, stable string) {
	st := time.Now().Nanosecond()
	// insert data
	res, err := db.Exec("insert into " + stable +
		" values (now, 100, 'beijing', 10,  true, 'one', 123.456, 123.456)" +
		" (now+1s, 101, 'shanghai', 11, true, 'two', 789.123, 789.123)" +
		" (now+2s, 102, 'shenzhen', 12,  false, 'three', 456.789, 456.789)")

	checkErr(err)

	affectd, err := res.RowsAffected()
	checkErr(err)

	et := time.Now().Nanosecond()
	fmt.Printf("insert data result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1E9)
}

func select_data(db *sql.DB, stable string) {
	st := time.Now().Nanosecond()

	rows, err := db.Query("select * from ? ", stable) // go text mode
	checkErr(err)

	fmt.Printf("%10s%s%8s %5s %9s%s %s %8s%s %7s%s %8s%s %4s%s %5s%s\n", " ", "ts", " ", "id", " ", "name", " ", "len", " ", "flag", " ", "notes", " ", "fv", " ", " ", "dv")
	var affectd int
	for rows.Next() {
		var ts string
		var name string
		var id int
		var len int8
		var flag bool
		var notes string
		var fv float32
		var dv float64

		err = rows.Scan(&ts, &id, &name, &len, &flag, &notes, &fv, &dv)
		checkErr(err)

		fmt.Printf("%s\t", ts)
		fmt.Printf("%d\t", id)
		fmt.Printf("%10s\t", name)
		fmt.Printf("%d\t", len)
		fmt.Printf("%t\t", flag)
		fmt.Printf("%s\t", notes)
		fmt.Printf("%06.3f\t", fv)
		fmt.Printf("%09.6f\n", dv)

		affectd++
	}

	et := time.Now().Nanosecond()
	fmt.Printf("insert data result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1E9)
	fmt.Printf("insert data result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1E9)
}

func drop_database_stmt(db *sql.DB, dbname string) {
	st := time.Now().Nanosecond()
	// drop test db
	stmt, err := db.Prepare("drop database ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(dbname)
	checkErr(err)

	affectd, err := res.RowsAffected()
	checkErr(err)

	et := time.Now().Nanosecond()
	fmt.Printf("drop database result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1E9)
}

func create_database_stmt(db *sql.DB, dbname string) {
	st := time.Now().Nanosecond()
	// create database
	//var stmt interface{}
	stmt, err := db.Prepare("create database ?")
	checkErr(err)

	//var res driver.Result
	res, err := stmt.Exec(dbname)
	checkErr(err)

	//fmt.Printf("Query OK, %d row(s) affected()", res.RowsAffected())
	affectd, err := res.RowsAffected()
	checkErr(err)

	et := time.Now().Nanosecond()
	fmt.Printf("create database result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1E9)
}

func use_database_stmt(db *sql.DB, dbname string) {
	st := time.Now().Nanosecond()
	// create database
	//var stmt interface{}
	stmt, err := db.Prepare("use " + dbname)
	checkErr(err)

	res, err := stmt.Exec()
	checkErr(err)

	affectd, err := res.RowsAffected()
	checkErr(err)

	et := time.Now().Nanosecond()
	fmt.Printf("use database result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1E9)
}

func insert_data_stmt(db *sql.DB, stable string) {
	st := time.Now().Nanosecond()
	// insert data into table
	stmt, err := db.Prepare("insert into ? values(?, ?, ?, ?, ?, ?, ?, ?) (?, ?, ?, ?, ?, ?, ?, ?) (?, ?, ?, ?, ?, ?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec(stable, "now", 1000, "'haidian'", 6, true, "'AI world'", 6987.654, 321.987,
		"now+1s", 1001, "'changyang'", 7, false, "'DeepMode'", 12356.456, 128634.456,
		"now+2s", 1002, "'chuangping'", 8, true, "'database'", 3879.456, 65433478.456)
	checkErr(err)

	affectd, err := res.RowsAffected()
	checkErr(err)

	et := time.Now().Nanosecond()
	fmt.Printf("insert data result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1E9)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
