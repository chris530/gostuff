package main

import (
	"strings"
	"io"
	"fmt"
)


func ReadAll(readers []io.Reader) (string, error) {
  var sb strings.Builder
  for _, r := range readers {
    _, err := io.Copy(&sb, r)
    if err != nil {
      return "", err
    }
    sb.WriteString(": ")
  }
  return sb.String(), nil
}


func main() {

  var err error
  var out string

  one := strings.NewReader("Hello, world!")
  two := strings.NewReader("How are you?")

  willWork := []io.Reader{one, two}

  out, err = ReadAll(willWork)
  
  if err != nil {
    panic(err)
  }

  fmt.Println(out)
	
}
