/**
* @author Striker007
*/
package config

import (
    "fmt"
    "os"
    "io/ioutil"
    // "parser"
    "strings"
)

func Read(fName string) (line string) {

    if len(fName) > 0 {

        f, fErr := os.Open(fName)
        data, fErr := ioutil.ReadFile(fName)
        if fErr != nil {
            fmt.Fprintf(os.Stderr,  "config : %v\n", fErr)
        }

        for lineNum, line := range strings.Split(string(data), "\n") {
            switch lineNum {
            // right now we interested only in 1st line
            case 0:
                continue
            case 1:
                return line
            }
        }
        f.Close()
    }
    return ""
}
