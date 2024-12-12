/*
	If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.

If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone.
 (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)

If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []int {
  content, err := os.ReadFile("./example2")

  if err != nil {
    panic(err)
  }

  strArr := strings.Split(strings.TrimSpace((string(content))), " ")
  intArr := make([]int, len(strArr))

  for i, str := range strArr {
    num, _ := strconv.ParseInt(str, 10, 64)
    intArr[i] = int(num)
  }

  return intArr
}

func yeah(hash map[int][]int, num int, idx int, blinking int) int {
  fmt.Println("yeah params: num", num)
  if idx >= blinking {
    return 0
  }

  cachedNum, exist := hash[num]
  if exist {
    if len(cachedNum) < blinking-idx {
    } else {
      return cachedNum[blinking - idx]
    }
  }

  var value int

  if num == 0 {
    value = yeah(hash, 1, idx+1, blinking)
    hash[num] = value + 1
    // fmt.Println("num", num,"value",value,"idx",idx)
    return hash[num]
  }

  tmpArr := make([]int, 0)
  tmp := num

  for tmp > 0 {
    oneItemArr := []int{tmp%10}
    tmpArr = append(oneItemArr, tmpArr...)
    tmp /= 10
  }

  if len(tmpArr) % 2 == 0 {
    a := 0
    b := 0
    halfIdx := len(tmpArr) / 2

    for i := 0; i < halfIdx; i++ {
      a *= 10 
      b *= 10

      a += tmpArr[i]
      b += tmpArr[halfIdx + i]
    }

    // fmt.Println("num", num,"value",value,"idx",idx, "a", a, "b", b, "tmpArr", tmpArr)
    value = yeah(hash, a, idx + 1, blinking) + yeah(hash, b, idx + 1, blinking)
    hash[num] = value + 2
    return hash[num]
  } else {
    value = yeah(hash, num * 2024, idx + 1, blinking)
    hash[num] = value + 1
    // fmt.Println("num", num,"value",value,"idx",idx)
    return hash[num]
  }
}

func main() {
  stones := readInput()
  hash := make(map[int]int)
  blinking := 4
  var answer int = 0

  for i := 0; i < len(stones); i++ {
    answer += yeah(hash, stones[i], 0, blinking)
  }

  fmt.Println(hash)
  fmt.Println(answer)
}
