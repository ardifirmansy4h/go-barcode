package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	linked := "https://www.linkedin.com/in/ardi-firmansyah-29346a22b"
	qrCode, err := qr.Encode(linked, qr.M, qr.Auto)
	if err != nil {
		fmt.Println("Error")
	}
	qrCode, _ = barcode.Scale(qrCode, 300, 300)
	file, _ := os.Create("file/qrlinkedd.png")
	defer file.Close()
	png.Encode(file, qrCode)
	fmt.Println("Barcode Created")
}
