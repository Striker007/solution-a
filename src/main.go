/**
* main.go - entry point
* @author Striker007
*/
package main

import (
        "fmt"
//        "config"
        "database/sql"
)

import _ "github.com/go-sql-driver/mysql"

func main() {
    // TODO config module
    db, err := sql.Open("mysql", "root:root@/adservone")
    rows, err := db.Query("SELECT id FROM bn_ads WHERE 1 LIMIT 1")
    defer rows.Close()

    for rows.Next() {
        var id int
        err = rows.Scan(&id)
        fmt.Print(id)
    }
     fmt.Print(err);
//    err = rows.Err() // get any error encountered during iteration
}
