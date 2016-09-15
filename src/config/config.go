
package config

   type T struct {
          A string
          B struct {
               RenamedC int `yaml:"c"`
               D        []int `yaml:",flow"`
          }
   }
