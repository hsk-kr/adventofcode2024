/**
1428664 wrong answer
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

func readInput() []string {
  content, err := os.ReadFile("./input")

  if err != nil {
    panic(err)
  }

  return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func yeah(r rune,m []string, visited map[position]bool, pos position, mArea map[position]bool, mPerimeter map[position]bool) {
  x := pos.x
  y := pos.y

  if x < 0 || y < 0 || x >= len(m[1]) || y >= len(m) || visited[pos] || rune(m[y][x]) != r {
    return
  }

  visited[pos] = true
  mArea[position{ x, y }] = true
  mPerimeter[position{ x, y }] = !(mArea[position{ x: x - 1, y: y }] && mArea[position{ x: x - 1, y: y - 1 }] && mArea[position{ x: x, y: y - 1}])
  mPerimeter[position{ x: x, y: y + 1 }] = !(mArea[position{ x: x - 1, y: y }] && mArea[position{ x: x - 1, y: y + 1 }] && mArea[position{ x: x, y: y + 1}])
  mPerimeter[position{ x: x + 1, y: y }] = !(mArea[position{ x: x + 1, y: y }] && mArea[position{ x: x + 1, y: y - 1 }] && mArea[position{ x: x, y: y - 1}])
  mPerimeter[position{ x: x + 1, y: y + 1 }] = !(mArea[position{ x: x + 1, y: y }] && mArea[position{ x: x + 1, y: y + 1 }] && mArea[position{ x: x, y: y + 1}])

  yeah(r, m, visited, position{ x: x - 1, y: y}, mArea, mPerimeter)
  yeah(r, m, visited, position{ x: x + 1, y: y}, mArea, mPerimeter)
  yeah(r, m, visited, position{ x: x, y: y - 1}, mArea, mPerimeter)
  yeah(r, m, visited, position{ x: x, y: y + 1}, mArea, mPerimeter)
}

func main() {
  m := readInput()
  visited := make(map[position]bool)

  answer := 0
  for y, row := range m {
    for x, r := range row {
      mPerimeter := make(map[position]bool)
      mArea := make(map[position]bool)

      if visited[position{x, y}] {
        continue
      }
      yeah(r, m, visited, position{ x, y }, mArea, mPerimeter)

      // if len(mArea) > 0 {
      //   fmt.Println(r, mArea, "/", mPerimeter)
      // }

      numArea := len(mArea)
      numPerimeter := 0
      
      for _, v := range mPerimeter {
        if v {
          numPerimeter++
        }
      }

      if (numArea != 0) {
        answer += numArea * numPerimeter
        fmt.Printf("%c %d x %d = %d / answer = %d\n", r, numArea, numPerimeter, numArea * numPerimeter, answer)
      }
    }
  }

  fmt.Println(answer)
}
