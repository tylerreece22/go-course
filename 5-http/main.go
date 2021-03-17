package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/* Custom Writer struct
Implicitly implements Writer interface
because it has a receiver function with
the correct function signature
 */
type logWriter struct {}

func main() {
	resp, error := http.Get("http://google.com")
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	fmt.Println(resp)


	// Manual way to print from response
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	// Go native approach
	//io.Copy(os.Stdout, resp.Body)len(bs)

	// Use custom type approach
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

// Garbage in -- Garbage out issue with Go interfaces
//func (logWriter) Write(bs []byte) (int, error) {
//	return 1, nil
//}


func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))

	return len(bs), nil
}