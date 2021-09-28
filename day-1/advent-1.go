package main

import (
  "fmt"
  "os"
  "log"
  "strings"
  "strconv"
)

func main() {

  // load file
  file, err := os.ReadFile("./data.txt")

  // handle error
  if err != nil {
    log.Fatal(err)
  }

  // store each line in array
  array := strings.Split(string(file),"\n")

  // trash empty last line
  array = array[:len(array)-1]

  // initialize freq
  freq := 0

  //initialize map to keep track of freq values
  freqmap := make(map[int]int)

  // for each array entry, perform the sign operation in front of the freq shift, by the freq shift and adjust freq
loop:
  for i := 0; i < len(array); i++{
      if string(array[i][0]) == "+" {

        shift, err := strconv.Atoi(array[i][1:])
        freq += shift

        if err != nil {
          log.Fatal(err)
        }

      } else {

        shift, err := strconv.Atoi(array[i][1:])
        freq -= shift

        if err != nil {
          log.Fatal(err)
        }
      }

      // if seen freq before, break and return duplicate
      if freqmap[freq] == 1 {
        fmt.Print(freq)
        break;
      } else {
        freqmap[freq] = 1
      }
      if i == len(array) - 1 {
        // couldve just used functions instead, but this is interesting! xD
        goto loop
      }
    }
}
