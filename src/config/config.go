/**
* @author Striker007
*/
package config

import (
    // "fmt"
    // "os"
    // "io/ioutil"
    "yaml_parser"
    // "strings"
    "errors"
)

var params map[string]interface{}

func Init(configFile string) (error) {

    if len(configFile) > 0 {

        f, fErr := os.Open(configFile)
        data, fErr := ioutil.ReadFile(configFile)
        if fErr != nil {
            fmt.Fprintf(os.Stderr,  "config : %v\n", fErr)
        }

        for lineNum, line := range strings.Split(string(data), "\n") {
            switch lineNum {
            // right now we interested only in 1st line
            case 0:
                continue
            case 1:
                tableName = line
                return true
            }
        }
        f.Close()
    }
    return false
}

func IsEmpty() {
    cfg, err := config.Get("unknown_param")
    if err != nil {
        fmt.Printf("config err: \n%v\n", err)
    }
}

func Get(paramName string) (string, error) {
    if len(params) == 0  {
        return "", errors.New("parameter list is empty, please Initialize config")
    }
    return "a", nil
}
