/**
* main.go - entry point
* @author Striker007
*/
package main

import (
	"fmt"
    _ "mysql"
    "github.com/icrowley/fake"
    "gopkg.in/yaml.v2"
)

var data = `
tables:
    test:
      LastName: LastName
      FirstName: FirstName
      gender: Gender
`

func main() {
    // TODO config module
    // mysql.ConnectToMysql();]

    fmt.Println(fake.FirstName())

    var tables interface{}

    err := yaml.Unmarshal([]byte(data), &tables)
    if err != nil {
            fmt.Printf("error: %v", err)
    }

    fmt.Printf("--- table:\n%v\n\n", tables.(map[interface {}]interface {}))

}
