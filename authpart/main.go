package main

import (
	"fmt"
	"math/rand/v2"
	"net"
	"net/http"
	"os/exec"
	"sync"
	"time"
)

func ipAddress() []net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}

	var ip net.IP
	var ipv4 []net.IP
	for _, addr := range addrs {
		switch v := addr.(type) {
		case *net.IPAddr:
			ip = v.IP
		case *net.IPNet:
			ip = v.IP
		}
		if ip != nil && ip.To4() != nil {
			ipv4 = append(ipv4, ip)
		}
	}
	return ipv4
}

func uuidGen() string {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("%s", uuid)
}

func OtpGenerate() int {
	return rand.IntN(1000000)
}

var (
	lastGenerated time.Time
	session       time.Time
	otp           int
)

func sessionHandler() string {
	return time.Now().Format("15:05:04")
}

var count int
var mu sync.Mutex

func counter() {
	mu.Lock()
	count++
	mu.Unlock()
}

func httpHandlers() {
	ipaddress := ipAddress()
	str := fmt.Sprintf("%s", ipaddress[1])
	http.HandleFunc("/otp", func(w http.ResponseWriter, r *http.Request) {
		if time.Since(lastGenerated) >= 30*time.Second {
			otp = OtpGenerate()
			lastGenerated = time.Now()
		}
		counter()
		fmt.Fprintf(w, "My otp for this session %v is: %06d\n", sessionHandler(), otp)
	})
	http.HandleFunc("/uuid", func(w http.ResponseWriter, r *http.Request) {
		uuid := uuidGen()
		counter()
		fmt.Fprintf(w, "My uuid for this session %v is: %s\n", sessionHandler(), uuid)
	})
	http.HandleFunc("/ipaddress", func(w http.ResponseWriter, r *http.Request) {
		counter()
		fmt.Fprintf(w, "My IPaddress for this session %v is: %v\n", sessionHandler(), str)
	})
	http.HandleFunc("/counter", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My counter is: %v\n", count)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("%v", err)
	}
}

func main() {
	httpHandlers()
}
