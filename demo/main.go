package main

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"os/exec"
	"os"
)

func isExistFile(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
   	}
	return false
}

func printQRCode(printer, filename string) {
	cmd := exec.Command("lpr","-P", printer, "-o", "landscape", "-o", "fit-to-page", "-o", "media=A4", filename)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}

func main() {
	data := "Do Huu Tai"
	qr_png := "/home/dotai/workspace/qr-test/qr.png"
	printer := "T58K"
	qrcode.WriteFile(data, qrcode.Medium, 256, qr_png)
	if isExistFile(qr_png) {
		printQRCode(printer, qr_png)
	} else {
		fmt.Println("QR Code is not exist")
	}
}
