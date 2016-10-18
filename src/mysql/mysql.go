/**
* main.go - entry point
* @author Striker007
*/
package mysql

import (
"fmt"
"database/sql"
)

import _ "github.com/go-sql-driver/mysql"

func ConnectToMysql() {

    // TODO config module

    dsn := "root:root@(127.0.0.1:3306)/test"
    fmt.Println("Connect by - " + dsn)
    db, err := sql.Open("mysql", dsn)

    if (err != nil) {
    	fmt.Println(err)
    }

    rows, err := db.Query("SELECT FirstName FROM test WHERE 1 LIMIT 1")
    defer rows.Close()
    
    if (err != nil) {
    	fmt.Println(err)
    }

    if(err == nil) {
    	for rows.Next() {
    		var FirstName string
    		_ = rows.Scan(&FirstName)
    		fmt.Println(FirstName)
    	}
    }
}
