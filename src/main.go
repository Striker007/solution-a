/**
* main.go - entry point
* @author Striker007
*/
package main

import (
    "config"
    "solutionA"
)

func main() {
    // TODO

    // get  command line args

    // init app configs
    // ... think about smth like config.getTableName
    tableName := config.Read("/home/vagrant/dumps/solution-a/config/tables.yaml")

    // connect to db
    // mysql.ConnectToMysql();

    // doing some app action
    if "" != tableName {
        solutionA.Fill(tableName)
        // like solution-a.job
            // fake 1 table
                // get table schema from config
                // generate fake data
                // update DB
            // generate yaml for fake'ing
                // read db schema
                // write config
    }

    // return exit code OR some info
}
