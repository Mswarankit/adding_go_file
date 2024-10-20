package main

import (
	"AddingGoFile/captcha/generate"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "captcha_test.html")
	})
	http.HandleFunc("/captcha", generate.GenerateCaptchaHandler)
	// http.HandleFunc("/verify", generate.VerifyCaptchaHandler)
	http.HandleFunc("/welcome", generate.WelcomeHandler)

	fmt.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
