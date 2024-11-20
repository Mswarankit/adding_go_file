package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8090/cointoss")
	checkErr(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println("Response:", string(body))
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
