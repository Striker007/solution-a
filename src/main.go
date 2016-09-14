/**
* main.go - entry point
* @author Striker007
*/
package main

import (
        "fmt"
        "config"
)

func main() {
   // test struct
   c := config.Main {20, 16}
   fmt.Sprintf("(%d, %d)", c.x, c.y)
}

