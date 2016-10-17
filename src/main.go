/**
* main.go - entry point
* @author Striker007
*/
package main

import (
    "mysql"
    // "config"
)

func main() {

    // TODO config module
    mysql.ConnectToMysql();
}
