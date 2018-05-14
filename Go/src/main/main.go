package main

import ("fmt"
  "net/http"
)

func main(){
  fmt.Println("Test")
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w,"This is my localhost")
  })

  http.ListenAndServe(":8080",nil)

}
