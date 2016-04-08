package main

  import (
    "fmt"
    "sort"
  )



      type People []string
      //included the below methods to implement the Sort interface
      func (p People) Len() int { return len(p)}
      func (p People) Less(i, j int) bool {return p[i] < p[j]}
      func (p People) Swap(i, j int) {p[i], p[j] = p[j], p[i]}

      type Numbros []int

      func (n Numbros) Len()int { return len(n)}
      func (n Numbros) Less( a, b int) bool {return n[a] < n[b]}
      func (n Numbros) Swap(a, b int) {n[a], n[b] = n[b], n[a]}


      func main() {

      studyGroup := People{"Zeno", "John", "Al", "Jenny"}

      fmt.Println(studyGroup)
      sort.Sort(studyGroup)
      fmt.Println(studyGroup)

      s := People{"Zeno", "John", "Al", "Jenny"}

      fmt.Println(s)
      sort.Sort(s)
      fmt.Println(s)

      n := Numbros{7, 4, 8, 2, 9, 19, 12, 32, 3}

      fmt.Println(n)
      sort.Ints(n)
      fmt.Println(n)

      k := sort.Reverse(n)

      fmt.Println(sort.Reverse(sort.IntSlice(n)))
      fmt.Println(k)


      var name interface{} = "ham"
      num, ok := name.(int)
      if ok {
        fmt.Printf("%T\n", num)
      } else {
        fmt.Printf("This value is not an integer")
      }



    }
