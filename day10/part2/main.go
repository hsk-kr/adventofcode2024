/**
1494 failed
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
  x int
  y int
}

func readInput() [][]int {
  content, err := os.ReadFile("./input")

  if err != nil {
    panic(err)
  }

  lines := strings.Split(strings.TrimSpace(string(content)), "\n")

  topoMap := make([][]int, len(lines))

  for i, line := range lines {
    intArr := make([]int, len(line))

    for j, h := range line {
      if h == '.' {
        intArr[j] = -1
      } else {
        intArr[j] = int(h - '0')
      }
    }

    topoMap[i] = intArr
  }

  return topoMap
}

func reachable(currentHeight int, topoMap [][]int, x,y int) int {
  if x < 0 || y < 0 || y >= len(topoMap) || x >= len(topoMap[0]) {
    return 0
  }

  if topoMap[y][x] == -1 || topoMap[y][x] != currentHeight {
    return 0
  }

  if currentHeight == 9 {
    return 1
  }

  currentHeight += 1

  cnt := 0
  cnt += reachable(currentHeight, topoMap, x - 1, y)
  cnt += reachable(currentHeight, topoMap, x + 1, y)
  cnt += reachable(currentHeight, topoMap, x, y - 1)
  cnt += reachable(currentHeight, topoMap, x, y + 1)
  
  return cnt
}

func main() {
  topoMap := readInput()
  answer := 0

  for y, line := range topoMap {
    for x, h := range line {
      if h == 0 {
        cnt := reachable(0, topoMap, x, y)
        fmt.Println("cnt", cnt, y, x)
        answer += cnt 
      }
    }
  }

  fmt.Println(answer)
}
