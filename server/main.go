package main

// import statements
import (
  "fmt" //print functions
  "html" //formatting
  "log" //outprint to console
  "net/http" //for server
)


func main() {
  
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hi")
  })

  log.Fatal(http.ListenAndServe(":8081", nil))
}
