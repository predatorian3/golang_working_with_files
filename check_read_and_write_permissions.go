//
package main

import (
  "log"
  "os"
)

func main() {
  file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
  if err != nil {
    if os.IsPermission(err) {
      log.Println("Error: Write permission denied.")
    }
  }
  log.Println("Write Permissions granted!")
  file.Close()

  file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
  if err != nil {
    if os.IsPermission(err) {
      log.Println("Error: Read Permissions denied.")
    }
  }
  log.Println("Read permissions granted!")
  file.Close()
}
