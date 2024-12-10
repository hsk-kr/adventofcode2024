// first answer: 6309364681516
package main

import (
	"fmt"
	"os" 
  "strings")

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
    for i := initBack; i > front; i -= 2 {
      if blocks[i][0] != -1 && blocks[front][len(blocks[front])-2] == -1 && blocks[front][len(blocks[front]) - 1] > 0 && blocks[i][1] <= blocks[front][len(blocks[front]) - 1] {
        blocks[front][len(blocks[front]) - 1] -= blocks[i][1]

        blocks[front] = append(blocks[front][:len(blocks[front])-2] , blocks[i][0], blocks[i][1], blocks[front][len(blocks[front])-2], blocks[front][len(blocks[front])-1])
        blocks[i][0] = -1

        if blocks[front][len(blocks[front]) - 1] == 0 {
          break
        }
      }
    }

    front += 2
  }

  answer := 0
  idx := 0

  for _, block := range blocks {
    for i := 0; i < len(block); i += 2 {
      for j := 0; j < block[i + 1]; j++ {
        fmt.Printf("%d*%d ", block[i], idx)
        if block[i] != -1 {
          answer += block[i] * idx
        }
        idx++
      }
    }
  }

  fmt.Println(answer)
}
