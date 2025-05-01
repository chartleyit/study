package main

import (
  "fmt"
)

func add(a int, b int) int {
  return a + b
}

func sub(a int, b int) int {
  return a - b
}

func mult(a int, b int) int {
  return a * b
}

func div(a int, b int) int {
  return a / b
}

func remain(a int, b int) int {
  return a % b
}

func main() {
  fmt.Println("1 + 2 =", add(1,2))
  fmt.Println("3 - 1 =", sub(3,1))
  fmt.Println("4 * 5 =", mult(4,5))
}
