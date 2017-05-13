package main

import (
	"io"
	"net/http"
  "log"
  "fmt"
  "time"
)


func indexHandler(w http.ResponseWriter, r *http.Request)  {
  t1 := time.Now()
  io.WriteString(w, "hello, gochord!\n")
  t2 :=time.Now()
  fmt.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
}


func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
