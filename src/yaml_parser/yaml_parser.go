/**
* @author Striker007
*/
package yaml_parser

import (
    "fmt"
    "gopkg.in/yaml.v2"
)

func ParseYaml(yamlLine string) (parsedLine interface{}) {

    err := yaml.Unmarshal([]byte(yamlLine), &parsedLine)
    if err != nil {
            fmt.Printf("Parser error: %v", err)
    }
    return
}
