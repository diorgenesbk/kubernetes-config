package main
import (
	"fmt"	
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request){
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	
	fmt.Fprintf(w, "Hello Success! Name: %s, Age: %s.", name, age)
}