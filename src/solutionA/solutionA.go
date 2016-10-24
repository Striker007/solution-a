/**
* @author Striker007
*/
package solutionA

import (
    "errors"
    _ "github.com/icrowley/fake"
    "config"
    _ "yaml_parser"
    "fmt"
)

func initialize() (error) {

    // if line := config.Init(CONFIG) {
    //     fmt.Prinln(line)
    // }

    return errors.New("parameter list is empty, please Initialize config")

    // init connection  o db
    // db.ConnectToDb(*cfg);
    
    // checkedTableName := parser.Parse(tableName)
    // parser.getTableSchema(checkedTableName)
    // fakeGenerate(schema map[string]string)
    // fillTable(tableName)
    return nil
}

func getTableSchema(tableName string) {

}

func fakeGenerate(fakeType map[string]string) (interface{}) {
    return fake.FirstName()
}

func FillOne() (error) {
    if err := initialize(); err != nil {
        return err
    }
    return nil
}