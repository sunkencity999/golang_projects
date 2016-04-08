package main

 import (
   "fmt"
   "encoding/json"
 )

 type user struct {//remember to capitalize the names of the fields within struct, so that they are exported!
   Name string
   Password string
   Id int
 }

   func (u user) fullName() string { //this is a method, which is a function that's attached to an 'object' (model, struct)
     return u.Name
   }

   func main() {
     u1 := user{"Christopher Bradford", "passwd", 1}
     u2 := user{"Tom Flanagan", "passwd", 2}
     us, _ := json.Marshal(u1)

     fmt.Println(us)
     fmt.Println(string(us))
     fmt.Println(u1.fullName())
     fmt.Println(u2.fullName())


   }
