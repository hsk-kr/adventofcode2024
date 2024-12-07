/*
*
1.

a + b + c + d

prev
total
answer
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) []string {
  content, err := os.ReadFile(path)
  
  if err != nil {
    panic(err)
  }

  return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func main() {
  lines := readInput("./input")

  var answer int64 = 0

  for _, line := range lines {
    nums := strings.Split(line, " ")

    shouldBe, err := strconv.ParseInt(nums[0][:len(nums[0]) -1], 10, 64)
    if err != nil {
      panic(err)
    }

    factors := make([]int64, 0)

    for i := 1; i < len(nums); i++ { var convertedNum int64
      convertedNum, err = strconv.ParseInt(nums[i], 10, 64)
      if err != nil {
        panic(err)
      }

      factors = append(factors, convertedNum)
    }

    var f func (idx int, prevTotal int64, total int64) bool;

    f = func (idx int, prevTotal int64, total int64) bool {
      if idx == len(factors) {
        return total == shouldBe
      }
      option := factors[idx]

      prevTotal = total
      total += option
      if f(idx + 1, prevTotal, total) {
        return true
      }

      total = prevTotal
      total *= option
      return f(idx + 1, prevTotal, total) 
    }

    if f(0, 0, 0) { 
      answer += shouldBe
    }
  }

  fmt.Println(answer)
}