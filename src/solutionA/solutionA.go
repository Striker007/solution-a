/**
* @author Striker007
*/
package solutionA

import (
    "fmt"
    "parser"
    _ "mysql"
    _ "github.com/icrowley/fake"
)

func Fill(table string) {
    yamlLine := parser.Parse(table)
    fmt.Printf("%v\n", yamlLine)
}