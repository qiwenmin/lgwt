package main

import (
	"fmt"
	"io"

	//"net/http"
	"os"
)

// Greet writes the greeting message via the writer.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!\n", name)
}

// MyGreeterHandler responses the greeting message on http.
// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
//     Greet(w, "world")
// }

func main() {
	Greet(os.Stdout, "world")
	//http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
