package main

import "fmt"

func lengthOfLongestSubstring(s string) (maxLength int) {
  seen := make(map[rune]int)

  start := -1

  for index, r := range s {
    idx, saw := seen[r]

    if !saw {
      seen[r] = index

      if start == -1 {
        start = index
      }

      if index == len(s)-1 && (index-start)+1 > maxLength {
        maxLength = (index-start)+1
      }
    } else {
      if index - start > maxLength {
        maxLength = index-start
      } else if idx == start && maxLength == 0 {
        maxLength = 1
      }

      for start <= idx {
        inRune := rune(s[start])
        delete(seen, inRune)
        start++
      }

      seen[r] = index
    }
    
  }

  return
}

func main() {
  fmt.Println(lengthOfLongestSubstring("ac"))
}