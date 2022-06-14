package images

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// imageGenerator es quien puede crear imagenes

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

// ImageAuxiliar crea una imagen en .png y lo guarda como archivo
// with the byte array received as parameter
func ImageAuxiliar(EncodeInformation []byte) error {
	//variables para los fors
	const a, b, c, d = 2, 20, 200, 2000
	// Imagen a color
	img := image.NewNRGBA(image.Rect(0, 0, defaultWidth, defaultHeight))
	//Se crea la imagen con determinado color, cada color es un rectangulo (16)
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
				G: uint8(int(EncodeInformation[1])),
				B: uint8(int(EncodeInformation[6])),
				A: 255,
			})
		}
		for x := c; x < d; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(int(EncodeInformation[15])),
				G: uint8(int(EncodeInformation[3])),
				B: uint8(int(EncodeInformation[5])),
				A: 255,
			})
		}
	}
	// se guarda la imagen en formato PNG
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	f, err := os.Create("NewAvatar.png")
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
		return err
	}
	// cierra el flujo
	if err := f.Close(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
