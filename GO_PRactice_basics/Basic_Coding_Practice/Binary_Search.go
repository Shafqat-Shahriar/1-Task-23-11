package main 

import (
  "fmt"
  "time"
)

// works when the array is sorted

func binarySearcher(arr []int, x int) int {
  i:= 0
  j:= len(arr)
  for i != j {
    var m = (i+j)/2
    if x == arr[m] {
      return m
    }
}
