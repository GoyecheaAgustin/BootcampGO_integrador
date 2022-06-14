package avatar

import (
	"avatar/pkg/avatar/encoder"
	"avatar/pkg/avatar/images"
	"log"
)

// cryptoEncoder codifica la informacion
type CryptoEncoder interface {
	EncodeInformation(strInformation string) (EncodeInformation []byte, err error)
}

// imageGenerator es quien puede generar imagenes
type ImageGenerator interface {
	BuildAndSaveImage(EncodeInformation []byte) error
}

// service contains funcionalities related to avatar generation
type Service struct {
	encoder   CryptoEncoder
	generator ImageGenerator
}

//nuevo servicio que ofrece el encoder y el generador
func AvatarGenerator() *Service {
	return &Service{
		encoder:   &encoder.CryptoEncoder{},
		generator: &images.ImageGenerator{},
	}
}

// Contiene informacion en este caso el nombre de usuario
type Information struct {
	UserName string
}

// GenerateAndSaveAvatar genera y guarda el avatar
func (s *Service) GenerateAndSaveAvatar(information Information) error {
	// defer para recuperarse de errores al encodear
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	// encodea la informacion
	encodeBytes, err := s.encoder.EncodeInformation(information.UserName)
	if err != nil {
		return err
	}
	// defer para recuperarse de errores al generar y guardar la imagen
	defer func() {
		if errImage := recover(); errImage != nil {
			log.Println(errImage)
		}
	}()
	// construye y guarda el png
	errImage := s.generator.BuildAndSaveImage(encodeBytes)
	if errImage != nil {
		return errImage
	}
	println("The avatar has been generated successfully")
	return nil
}
