package main

import (
	"errors"
	"fmt"
	"image/png"

	//"log"
	"os"
	"os/exec"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
)

// TODO: Figure out how to trim the filepath so the barcode can be named and saved accordingly
func generateBarcode(imageFilepath string, barcodeFilepath string) (string, error) {
	writer := oned.NewCode128Writer()

	img, err := writer.Encode(imageFilepath, gozxing.BarcodeFormat_CODE_128, 250, 50, nil)
	if err != nil {
		return "", errors.New("invalid input file")
	}

	file, err := os.Create(barcodeFilepath)
	if err != nil {
		return "", errors.New("could not generate new file")
	}

	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return "", errors.New("could not encode barcode as PNG")
	}

	return barcodeFilepath, nil

}

// TODO: Implement barcode scanning function


// TODO: Implement main function
func main() {
	/*
	_, err := generateBarcode("shirt_print_image.prn", "print_image_barcode.png")
	if err != nil {
		log.Fatal("could not generate barcode")
	}
	path, err := scanBarcode("print_image_barcode.png")
	if err != nil {
		log.Fatal("invalid barcode path")
	}
	fmt.Println(path.GetText())

	cmd := exec.Command("code", path.GetText())
	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	*/
	fmt.Println("Input Barcode:")
	var path string
	fmt.Scanln(&path)


	cmd := exec.Command("lp", path)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

}
