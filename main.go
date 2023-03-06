package main

import (
  "fmt"
  //"gihub.com/Toor3-14/algorithms/abc"
  "gihub.com/Toor3-14/algorithms/sort"
  "gihub.com/Toor3-14/algorithms/search"
)


func main() {
  x := []string{"Bob", "Jaque", "Jack", "Chris", "Johne", "Nocola"}
  sort.BubleCustom(x)
  fmt.Println(x)
  fmt.Println(search.BinaryCustom(x,"Bob"))
}

