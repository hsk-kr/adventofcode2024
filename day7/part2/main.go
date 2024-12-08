/*
*
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
  sna := strconv.FormatUint(na, 10)
  snb := strconv.FormatUint(nb, 10)

  num, err := strconv.ParseUint(sna + snb, 10, 64)

  if err != nil {
    panic(err)
  }

  return num
  // if na == 0 {
  //   ///fmt.Printf("returns %d - na: %d, nb: %d\n", nb, na, nb);
  //   return nb
  // }
  //
  // var numToMultiply uint64 = 1
  // tmpN := nb
  //
  // for tmpN > 0 {
  //   tmpN /= 10
  //   numToMultiply *= 10
  // }
  //
  // return na * numToMultiply + nb
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

    // fmt.Println(factors, nums)
    // if len(nums) <= 2 {
    //   fmt.Println("fuck", nums)
    // }
    // if shouldBe <= 0 {
    //   fmt.Println("shouldBe is less than or equal to zero", shouldBe)
    // }
    // if len(factors)<= 0 {
    //   fmt.Println("there is no factors", factors, nums)
    // }

    var f func (idx int, total uint64, test string) (bool, string)

    f = func (idx int, total uint64, test string) (bool, string) {
      if total > shouldBe { 
        return false, ""
      }

      if len(factors) == 1 {
        return factors[0] == shouldBe, test
      }
      if idx == len(factors)  {
        return total == shouldBe, test
      }

      option := factors[idx]
      prevTotal := total

      // if option == 0 {
      //   fmt.Println("option is zero", nums)
      // }

      var rst bool
      prevTest := test
      total = prevTotal + option
      if rst, test = f(idx + 1, total, test+"+"); rst {
        return true, test
      }

      test = prevTest
      total = prevTotal * option
      if rst, test = f(idx + 1, total, test+"*"); rst  {
        return true, test
      }

      test = prevTest
      total = concatenate(prevTotal, option)
      return f(idx + 1, total, test+"|")
    }

    var rst bool
    var test string
    if rst, test = f(0, 0, test); rst { 
      var testOutput string = ""

      for i, factor := range factors {
        if len(factors) - 1 == i {
          testOutput = fmt.Sprintf("%s %d", testOutput, factor)
        } else {
          testOutput = fmt.Sprintf("%s %d %c", testOutput, factor, test[i + 1])
        }
      }

      fmt.Println(answer, shouldBe, testOutput)
      answer += shouldBe
    }
  }

  fmt.Println(answer)
}
