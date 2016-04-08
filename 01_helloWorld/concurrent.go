package main

  import "fmt"

  func Lister(users ...string) chan string{
      names := make(chan string)
       go func(){
         for _, i := range users {
           names <- i
         }
       }()

       go func(){
         for _, i := range users {
           names <- i + " User"
           close(names)

         }
       }()

       return names
  }

    func main(){
    first := Lister(Alan, Simon, Theodore, Christopher)
    for n := range first {
    fmt.Println(n)
  }
    }
