
package config

import "testing"
//import "fmt"
//import "gopkg.in/yaml.v2"


// sample
//        var Data = `
//        a: Easy!
//        b:
//          c: 2
//          d: [3,4]
//        `

//   type Config struct {
//          A string
//          B struct {
//               RenamedC int `yaml:"c"`
//               D        []int `yaml:",flow"`
//          }
//   }

func TestConfig(t *testing.T) {

  c := Config{}

  if (c == Config{}) {
      t.Errorf("Config expected creating, actual")
  }
    /**    c := Config{
            a: "Easy!",
            b: {
                c: 2,
                d: {3,4}
            }
        }
**/


// https://github.com/go-yaml/yaml/tree/v2
          

//        t := T{}

//        err := yaml.Unmarshal([]byte(config.Data), &t)
//        if err != nil {
//                log.Fatalf("error: %v", err)
//        }
//        fmt.Printf("--- t:\n%v\n\n", t)
//
//
//        d, err := yaml.Marshal(&t)
//        if err != nil {
//                log.Fatalf("error: %v", err)
//        }
//        fmt.Printf("--- t dump:\n%s\n\n", string(d))
//
//
//        m := make(map[interface{}]interface{})
//
//        err = yaml.Unmarshal([]byte(config.Data), &m)
//        if err != nil {
//                log.Fatalf("error: %v", err)
//        }
//        fmt.Printf("--- m:\n%v\n\n", m)
//
//
//        d, err = yaml.Marshal(&m)
//        if err != nil {
//                log.Fatalf("error: %v", err)
//        }
//        fmt.Printf("--- m dump:\n%s\n\n", string(d))

}


