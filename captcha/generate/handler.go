package generate

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

var captchaStore = NewCaptchaStore()

func GenerateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	captchaText := generateCaptcha(6)
	id := generateCaptcha(12)
	captchaStore.Set(id, captchaText)

	var buf bytes.Buffer
	if err := createCaptchaImage(&buf, captchaText); err != nil {
		http.Error(w, "Failed to generate CAPTCHA", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"captchaId": "%s", "captchaImage": "data:image/png;base64,%s"}`,
		id, base64.StdEncoding.EncodeToString(buf.Bytes()))
}

func VerifyCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	captchaId := r.FormValue("captchaId")
	userInput := r.FormValue("captchInput")

	storedValue, exists := captchaStore.Get(captchaId)
	if !exists {
		http.Error(w, "CAPTCHA expired or not found", http.StatusBadRequest)
		return
	}
	if storedValue == userInput {
		captchaStore.Delete(captchaId)
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	} else {
		http.Error(w, "Incorrect CAPTCHA", http.StatusBadRequest)
	}
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Site and Time %s", time.Now().Format(time.RFC3339))
}
