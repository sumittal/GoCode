package main

import (
  "fmt"
)

func main() {
  out := comma("12345678")
  fmt.Println("output : ", out)
}
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
fmt.Println("Debug - ", s)
  n := len(s)
  if n <= 3 {
	return s
  }
  return comma(s[:n-3]) + "," + s[n-3:]
}
