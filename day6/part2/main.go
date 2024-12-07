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

type positionForce struct {
  x int
  y int
}

type inputInfo struct {
  visited map[position]bool
  obstacles map[position]bool
}


const (
  UP = iota
  RIGHT
  DOWN
  LEFT
)

func getPositionForceFromDir(dir int) position {
  x := 0
  y := 0

  if dir == UP {
    y -= 1
  } else if dir == RIGHT {
    x += 1
  } else if dir == DOWN {
    y += 1
  } else if dir == LEFT {
    x -= 1
  }

  return position {
    x,
    y,
  }
}

func readInput(path string) []string {
  buf, err := os.ReadFile(path)

  if err != nil {
    panic(err)
  }

  lines := strings.Split(strings.TrimSpace(string(buf)), "\n")

  return lines
}

func getStartPoint(pathLines []string) position {
  for y, pathLine := range pathLines {
    for x, char := range pathLine {
      if char == '^' {
        return position{
          x,
          y,
        }
      }
    }
  }

  panic("fail to find start point")
}

func isOutOfRange(pathLines []string, pos position) bool {
  return pos.x < 0 || pos.y < 0 || pos.y >= len(pathLines) || pos.x >= len(pathLines[0])
}

func nextDir(dir int) int {
  if dir < 0 || dir > 3 {
    panic(fmt.Sprintf("wrong direction: %d", dir))
  }
 
  return (dir + 1) % 4
}

func getAllVisitiedAndObstaclesPos(pathLines []string, startPosition position) *inputInfo {
  inputInfo := &inputInfo{
    visited: make(map[position]bool),
    obstacles: make(map[position]bool),
  }

  pos := startPosition
  prevPosition := position{x: 0, y: 0}
  dir := UP

  for !isOutOfRange(pathLines, pos) {
    if pathLines[pos.y][pos.x] == '#' {
      inputInfo.obstacles[pos] = true;
      pos = prevPosition
      dir = nextDir(dir)
      continue
    }

    inputInfo.visited[pos] = true

    force := getPositionForceFromDir(dir)

    prevPosition = pos
    pos.x += force.x
    pos.y += force.y
  }

  return inputInfo
}

func hasVisitiedWithSameDir(dirs []int, dir int) bool {
  for _, d := range dirs {
    if d == dir {
      return true
    }
  }

  return false
}

func canEscape(pathLines []string, startPosition position, obstacles map[position]bool) bool {
  pos := startPosition
  prevPosition := position{x: 0, y: 0}
  dir := UP
 
  visited := make(map[position]int)
  please := 0

  for !isOutOfRange(pathLines, pos) {
    if obstacles[pos] {
      pos = prevPosition
      dir = nextDir(dir)
      continue
    }

    force := getPositionForceFromDir(dir)

    visitedDir, exists := visited[pos]
    please++
    if exists && visitedDir == dir {
      fmt.Println(please, visitedDir, dir, pos, startPosition)
      return false
    }

    if pos != startPosition {
      visited[pos] = dir
    }else {
      fmt.Println("ignore first one")
    }

    prevPosition = pos
    pos.x += force.x
    pos.y += force.y

  }

  return true 
}

func main() {
  pathLines := readInput("./input")

  startPos := getStartPoint(pathLines)

  inputInfo := getAllVisitiedAndObstaclesPos(pathLines, startPos)

  answer := 0

  for pos := range inputInfo.visited {
    inputInfo.obstacles[pos] = true

    if !canEscape(pathLines, startPos, inputInfo.obstacles) {
      answer++
    }

    inputInfo.obstacles[pos] = false
  }
  
  fmt.Printf("Answer: %d\n", answer);
}

/*
1. Repeat until it's escape. Track, all the visited positions and all the obstacles.
2. Put an extra block all the visited position.
3. Run the first step and see if this is possible to escape. while it's moving, it keeps it's direction
   when at the same position and same direction, it means that it can't escape
4. Repeaet until trying all the possibiltieis of second one
*/
