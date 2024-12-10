// first answer: 6309364681516
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) string {
  content, err := os.ReadFile(path)

  if err != nil {
    panic(err)
  }

  return strings.TrimSpace(string(content))
}

func main() {
  diskMap := readInput("./input")

  blocks := make([][]int, 0)

  for i, strAmount := range diskMap {
    amount := int(strAmount - '0')

    pair := make([]int, 2)

    if i % 2 == 0 {
      pair[0] = i / 2
      pair[1] = amount
    } else {
      pair[0] = -1
      pair[1] = amount
    }
    
    blocks = append(blocks, pair)
  }

  initBack := len(blocks) - 1
  if len(blocks) % 2 == 0 {
    initBack = len(blocks) - 2
  }
  front := 1


  for front < initBack {
    found := true 

    for found { 
      found = false
      for i := initBack; i > front; i -= 2 {
        if blocks[i][1] > 0 && blocks[i][1] <= blocks[front][1] {
          blocks[front] = append(blocks[front], blocks[i][0], blocks[i][1])
          blocks[front][1] -= blocks[i][1]
          blocks[i][0] = -1
          initBack -= 2
          found = true
        }
      }
    }

    front += 2
  }

  for i := 1; i < len(blocks); i += 2 {
    block := blocks[i]
    if block[0] == -1 && block[1] > 0 {
      blocks[i] = append(block[2:], block[:2]...)
    }
  }

  final := ""

  for _, block := range blocks {
    for i := 0; i < len(block); i += 2 {
      char := strconv.Itoa(block[i])
      if char == "-1" {
        char = "."
      }
      for j := 0; j < block[i + 1]; j++ {
        final += char
      }
    }
  }

  fmt.Println(final)
  answer := 0

  for i, c := range final {
    if c != '.' {
      num := c - '0'
      answer += int(num) * i
    }
  }

  fmt.Println(answer)
}
