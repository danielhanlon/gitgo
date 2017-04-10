package main
import (
  "fmt"
  flag "github.com/ogier/pflag"
)

var (
  user string
)

func main() {
  fmt.Println("Hello, World")
  flag.Parse()
}

//Run before main
func init() {
  //expecting a string, variable to bind to, flag to expect, default value, description
  flag.StringVarP(&user, "user", "u", "", "Search Users")
}
