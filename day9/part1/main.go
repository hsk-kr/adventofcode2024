// first answer: 6309364681516
package main

import (
	"fmt"
	"os"
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

  back := len(blocks) - 1
  if len(blocks) % 2 == 0 {
    back = len(blocks) - 2
  }
  front := 1


  for front < back {
    if blocks[front][1] > blocks[back][1] {
      blocks[front] = append(blocks[front], blocks[back][0], blocks[back][1])
      blocks[front][1] -= blocks[back][1]
      blocks[back][1] = 0
      back -= 2
    } else if blocks[front][1] < blocks[back][1] {
      blocks[front] = append(blocks[front], blocks[back][0], blocks[front][1])
      blocks[back][1] -= blocks[front][1]
      blocks[front][1] = 0
      front += 2
    } else {
      blocks[front] = append(blocks[front], blocks[back][0], blocks[back][1])
      blocks[front][1] = 0
      blocks[back][1] = 0
      front += 2
      back -= 2
    }

  }

  answer := 0
  idx := 0

  for _, block := range blocks {

    for i := 0; i < len(block); i += 2 {
      if (block[i] == -1) {
        continue
      }
      for j := 0; j < block[i + 1]; j++ {
        answer += block[i] * idx;
        fmt.Printf("%d * %d = %d\n", block[i], idx, answer);
        idx++
      }
    }
  }

  fmt.Println(answer)
}
