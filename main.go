package main

import (
	"errors"
	"image"
	"image/png"
	"log"
	"os"

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
func scanBarcode(barcodeFilepath string) (*gozxing.Result, error) {
	file, err := os.Open(barcodeFilepath)
	if err != nil {
		return nil, errors.New("could not open barcode")
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, errors.New("could not decode barcode file")
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return nil, errors.New("could not generate binary bitmap from barcode image")
	}

	reader := oned.NewCode128Reader()
	result, err := reader.Decode(bmp, nil)
	if err != nil {
		return nil, errors.New("invalid barcode")
	}

	return result, nil
}

// TODO: Implement main function
func main() {
	_, err := generateBarcode("images/go.png", "images/barcode1.png")
	if err != nil {
		log.Fatal("could not generate barcode")
	}
	path, err := scanBarcode("./images/barcode1.png")
	if err != nil {
		log.Fatal("invalid barcode path")
	}
	print(path.GetText())
	

}
