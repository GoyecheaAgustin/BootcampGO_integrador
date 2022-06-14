package images

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// imageGenerator is someone who can generate an images

type ImageGenerator struct{}

func (g *ImageGenerator) BuildAndSaveImage(EncodeInformation []byte) error {
	err := ImageAuxiliar(EncodeInformation)
	if err != nil {
		return err
	} else {
		return nil
	}
}

const (
	defaultWidth  = 500
	defaultHeight = 500
)

// ImageAuxiliar generate an image in png format and save it in a file
// with the byte array received as parameter
func ImageAuxiliar(EncodeInformation []byte) error {
	//for's variables
	const a, b, c, d = 2, 20, 200, 2000
	// Create a colored image
	img := image.NewNRGBA(image.Rect(0, 0, defaultWidth, defaultHeight))
	// Fill the image with the given color. The image is filled in 16 rectangles
	for y := 0; y < a; y++ {
		for x := 0; x < a; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[0])),
				G: uint8(int(EncodeInformation[1])),
				B: uint8(int(EncodeInformation[2])),
				A: 255,
			})
		}
		for x := a; x < b; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[1])),
				G: uint8(int(EncodeInformation[2])),
				B: uint8(int(EncodeInformation[3])),
				A: 255,
			})
		}
		for x := b; x < c; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[2])),
				G: uint8(int(EncodeInformation[3])),
				B: uint8(int(EncodeInformation[4])),
				A: 255,
			})
		}
		for x := c; x < d; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[3])),
				G: uint8(int(EncodeInformation[4])),
				B: uint8(int(EncodeInformation[5])),
				A: 255,
			})
		}
	}
	for y := a; y < b; y++ {
		for x := 0; x < a; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[4])),
				G: uint8(int(EncodeInformation[5])),
				B: uint8(int(EncodeInformation[6])),
				A: 255,
			})
		}
		for x := a; x < b; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[5])),
				G: uint8(int(EncodeInformation[6])),
				B: uint8(int(EncodeInformation[7])),
				A: 255,
			})
		}
		for x := b; x < c; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[6])),
				G: uint8(int(EncodeInformation[7])),
				B: uint8(int(EncodeInformation[8])),
				A: 255,
			})
		}
		for x := c; x < d; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[7])),
				G: uint8(int(EncodeInformation[8])),
				B: uint8(int(EncodeInformation[9])),
				A: 255,
			})
		}
	}
	for y := b; y < c; y++ {
		for x := 0; x < a; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[8])),
				G: uint8(int(EncodeInformation[9])),
				B: uint8(int(EncodeInformation[10])),
				A: 255,
			})
		}
		for x := a; x < b; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[9])),
				G: uint8(int(EncodeInformation[10])),
				B: uint8(int(EncodeInformation[11])),
				A: 255,
			})
		}
		for x := b; x < c; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[10])),
				G: uint8(int(EncodeInformation[11])),
				B: uint8(int(EncodeInformation[12])),
				A: 255,
			})
		}
		for x := c; x < d; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[11])),
				G: uint8(int(EncodeInformation[12])),
				B: uint8(int(EncodeInformation[13])),
				A: 255,
			})
		}
	}
	for y := c; y < d; y++ {
		for x := 0; x < a; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[12])),
				G: uint8(int(EncodeInformation[13])),
				B: uint8(int(EncodeInformation[14])),
				A: 255,
			})
		}
		for x := a; x < b; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[13])),
				G: uint8(int(EncodeInformation[14])),
				B: uint8(int(EncodeInformation[15])),
				A: 255,
			})
		}
		for x := b; x < c; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[14])),
				G: uint8(int(EncodeInformation[15])),
				B: uint8(int(EncodeInformation[0])),
				A: 255,
			})
		}
		for x := c; x < d; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[15])),
				G: uint8(int(EncodeInformation[0])),
				B: uint8(int(EncodeInformation[1])),
				A: 255,
			})
		}
	}
	// Save the image to a PNG file.
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	f, err := os.Create("avatar.png")
	if err != nil {
		log.Fatal(err)
		return err
	}
	// Encode the image in PNG format.
	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
		return err
	}
	// Close the file.
	if err := f.Close(); err != nil {
		log.Fatal(err)
		return err
	}
	println("avatar created")
	return nil
}
