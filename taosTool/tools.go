package taosTool

import (
	"database/sql"
	"fmt"
	_ "github.com/satng/sensors-gateway-grpc/taosSql"
	"os"
	"time"
)

const (
	CONNDB = "%s:%s@/tcp(%s)/%s"
	DRIVER = "taosSql"
	DBHOST = "127.0.0.1"
	DBUSER = "root"
	DBPASS = "taosdata"
	DBNAME = "sensors_test"
)

var (
	globalDB *sql.DB
)

func InitDB() {
	// open connect to taos server
	connStr := fmt.Sprintf(CONNDB, DBUSER, DBPASS, DBHOST, DBNAME)
	db, err := sql.Open(DRIVER, connStr)
	if err != nil {
		fmt.Println("Open database error: %s\n", err)
		os.Exit(0)
	}
	globalDB = db

	_, err = globalDB.Exec("use " + DBNAME)
	checkErr(err)

	fmt.Println("Taos database ok")
}
func CloseDB() {
	globalDB.Close()
}
func Insert(sql string) {
	st := time.Now().Nanosecond()
	res, err := globalDB.Exec(sql)
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
