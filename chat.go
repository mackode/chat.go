package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  ai := NewAI()
  ai.init()
  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print("Ask: ")
    if !scanner.Scan() {
      break
    }
    ai.printResp(scanner.Text())
  }
}
