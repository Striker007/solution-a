/**
* @author Striker007
*/
package parser

import (
    "fmt"
    "gopkg.in/yaml.v2"
)

func Parse(yamlLine string) (parsedLine interface{}) {

    err := yaml.Unmarshal([]byte(yamlLine), &parsedLine)
    if err != nil {
            fmt.Printf("error: %v", err)
    }
    return
}
