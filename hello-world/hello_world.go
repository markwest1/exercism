package hello

import (
  "fmt"
)

const testVersion = 2

// HelloWorld prints "Hello name" or, if name is empty, "Hello World!", to
// stdout.
func HelloWorld(name string) string {
  object := "World"

  if name != "" {
    object = name
  }

  return fmt.Sprintf( "Hello, %s!", object)
}
