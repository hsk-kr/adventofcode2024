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
func readInput(path string) [][]string {
  content, err := os.ReadFile(path)
  
  if err != nil {
    panic(err)
  }

  result := make([][]string, 0)
  lines := strings.Split(strings.TrimSpace(string(content)), "\n")

  for _, line := range(lines) {
    result = append(result, strings.Split(line, ""))
  }

  return result
}

func getAntinodePositions(antennaMap [][]string, posA position, posB position) []position {
  posList := make([]position, 0)
  force := position {x: 0, y: 0}

  force.x = posA.x - posB.x
  force.y = posA.y - posB.y

  // same position
  if force.x == 0 && force.y == 0 {
    return posList
  }

  pos := position{x: posA.x, y: posA.y}

  for {
    pos.x = pos.x + force.x 
    pos.y = pos.y + force.y

    if pos.y >= len(antennaMap) || pos.x >= len(antennaMap[0]) || pos.y < 0 || pos.x < 0 {
      break
    }

    posList = append(posList, position { x: pos.x, y : pos. y })
  }

  return posList
  
}

func countAntinodes(antennaMap [][]string) int {
  count := 0

  for _, lines := range antennaMap {
    for _, a := range lines {
      if a != "." {
        count++
      }
    }
  }

  return count
}

func printMap(m [][]string) {
  for _, line := range(m) {
    fmt.Println(line)
  }
}

func main() {
  antennaMap := readInput("./input")
  antennaPosMap := make(map[string][]position)

  for y, line := range antennaMap{
    for x, antenna := range line {
      if antenna == "." {
        continue
      }

      pos := position{x, y}
      posList, exist := antennaPosMap[antenna]

      if !exist {
        posList = make([]position, 0)
      }

      posList = append(posList, pos)
      antennaPosMap[antenna] = posList
    }
  }

  for antenna, posList := range antennaPosMap {
    for _, posA := range posList {
      for _, posB := range posList {
        antinodePosList := getAntinodePositions(antennaMap, posA, posB)
        fmt.Println(antinodePosList)
        
        for _, antinodePos := range antinodePosList {
          if  !strings.Contains(antennaMap[antinodePos.y][antinodePos.x], antenna) {
            antennaMap[antinodePos.y][antinodePos.x] += antenna 
          }
        }
      }
    }
  }


  printMap(antennaMap)
  fmt.Println(countAntinodes(antennaMap))
}

/**
  find all antennas and keep the positions in the map
  key will be a characdter and values will be a position

  and iterate each antenna positions and n^2 iteration 
  i = 0 .. length -1
  j = 0 .. length -1
  antinode.x = pos[i].x - (pos[i].x - pos[j].x);
  antinode.y = pos[i].y - (pos[i].y - pos[j].y);

  if isRange(antinode) && map
    antinodePos[antinode] = true

  
  answer len(antinodePos)

  no it's wrong

  antinode locations may be able to be overlapped.

  try to store # + someother antenna
*/
