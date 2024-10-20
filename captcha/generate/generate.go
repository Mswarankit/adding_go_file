package generate

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/rand/v2"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func createCaptchaImage(w io.Writer, text string) error {
	width, height := 140, 80
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.White)
		}
	}
	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}

	// Create a larger font face
	face := truetype.NewFace(f, &truetype.Options{
		Size: 30, // Adjust this value to change the font size
		DPI:  72,
	})

	drawer := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: face,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(10 * 64), Y: fixed.Int26_6(50 * 64)},
	}
	drawer.DrawString(text)

	// Add noise
	for i := 0; i < 1000; i++ {
		x := rand.IntN(width)
		y := rand.IntN(height)
		img.Set(x, y, color.Black)
	}

	return png.Encode(w, img)
}

// func createCaptchaHandler(w http.ResponseWriter, r *http.Request) {
// 	captchaText := generateCaptcha(6)
// 	// Store captchaText in session or database for later verification

// 	var buf bytes.Buffer
// 	if err := createCaptchaImage(&buf, captchaText); err != nil {
// 		http.Error(w, "Failed to generate CAPTCHA", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "image/png")
// 	w.Write(buf.Bytes())
// }
