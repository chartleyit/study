package main

import "fmt"

func getLetters(s string) map[rune]int {
  letters := make(map[rune]int)
  

  for _, chr := range s {
    _, exists := letters[chr]
    if exists {
      letters[chr]++
    } else {
      letters[chr] = 1
    }
  }

  return letters
}

func RansomeNote(n string, m string) bool {
  makeNote := true

  available := getLetters(m)

  // TODO add note logic
  for _, e := range n {
    _, exists := available[e]
    if exists && available[e] > 0 {
      available[e]--
    } else {
      makeNote = false
      break
    }
  }

  return makeNote
}

func main() {
	magazine := "aabb"
	note := "aa"

  fmt.Println(RansomeNote(note, magazine))
}
