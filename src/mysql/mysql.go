/**
* main.go - entry point
* @author Striker007
*/
package mysql

import (
"os"
"fmt"
"database/sql"
)

import _ "github.com/go-sql-driver/mysql"

func ConnectToMysql() {

    // TODO config module

    dsn := os.Getenv("MYSQL_DSN") // root:my-secret-pw@/test
    fmt.Println(dsn)
    db, err := sql.Open("mysql", dsn)

    if (err != nil) {
    	fmt.Println(err)
    }

    rows, err := db.Query("SELECT id FROM test WHERE 1 LIMIT 1")
    defer rows.Close()
    
    if (err != nil) {
    	fmt.Println(err)
    }

    if(err == nil) {
    	for rows.Next() {
    		var id int
    		_ = rows.Scan(&id)
    		fmt.Print(id)
    	}
    }
}
