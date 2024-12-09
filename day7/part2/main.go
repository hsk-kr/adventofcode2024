/*
*zz
1.

a + b + c + d

prev
total
answer

part1: 5837374519342

5837374519342
492383931656533

first try: 492383931656533
after updating: 492383931656533
after refactoring concatenate: 492383931656533
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

func concatenate(na uint64, nb uint64) uint64 {
  if na == 0 {
    return nb
  }

  var numToMultiply uint64 = 1
  tmpN := nb

  for tmpN > 0 {
    tmpN /= 10
    numToMultiply *= 10
  }

  return na * numToMultiply + nb
}

func main() {
  lines := readInput("./input")

  var answer uint64 = 0

  for _, line := range lines {
    nums := strings.Split(line, " ")

    shouldBe, err := strconv.ParseUint(nums[0][:len(nums[0]) -1], 10, 64)
    if err != nil {
      panic(err)
    }

    factors := make([]uint64, 0)

    for i := 1; i < len(nums); i++ {
      var convertedNum uint64
      convertedNum, err = strconv.ParseUint(nums[i], 10, 64)
      if err != nil {
        panic(err)
      }

      factors = append(factors, convertedNum)
    }

    var f func (remain []uint64, total uint64) bool

    f = func (remain []uint64, total uint64) bool {
      if len (remain) == 0 { 
        return total == shouldBe
      }

      firstNum := remain[0]
      remain = remain[1:]
      if f(remain, total + firstNum) {
        return true
      }

      if f(remain, total * firstNum)  {
        return true
      }

      return f(remain, concatenate(total, firstNum))
    }

    firstNum := factors[0]
    factors = factors[1:]
    if f(factors, firstNum) { 
      answer += shouldBe
    }
  }

  fmt.Println(answer)
}
