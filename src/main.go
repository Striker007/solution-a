/**
* main.go - entry point
* @author Striker007
*/
package main

import (
    "fmt"
    "os"
    "solutionA"
)

const EXIT_NORMAL int = 0
const EXIT_ERR int = 6
const CONFIG string = "/home/vagrant/dumps/solution-a/config/tables.yaml"

func main() {

    helpMessage := "NAME solution-a \n fake data generator \n" +
    "SYNOPSIS \n solution-a [OPTION]...\n" +
    "DESCRIPTION \n fill your database with fake data \n" +
    "OPTIONS \n fillone - fill one table from config \n\n"

    args := os.Args[1:]
    for _, arg := range args {
        switch arg {        
        case "fillone":
            errA := solutionA.FillOne()
            if nil != errA {
                fmt.Printf("something going wrong: %v\n\n", errA)
                os.Exit(EXIT_ERR)
            }
        }          
    }

    fmt.Print(helpMessage)
    os.Exit(EXIT_ERR)
}
