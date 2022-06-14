package avatar

import (
	"avatar/pkg/avatar/encoder"
	"avatar/pkg/avatar/images"
	"log"
)

// cryptoEncoder is someone who can encode information in a cryptographic way
type CryptoEncoder interface {
	EncodeInformation(strInformation string) (EncodeInformation []byte, err error)
}

// imageGenerator is someone who can generate an images
type ImageGenerator interface {
	BuildAndSaveImage(EncodeInformation []byte) error
}

// service contains funcionalities related to avatar generation
type Service struct {
	encoder   CryptoEncoder
	generator ImageGenerator
}

// NewService creates a new service with the given encoder and generator
func AvatarGenerator() *Service {
	return &Service{
		encoder:   &encoder.CryptoEncoder{},
		generator: &images.ImageGenerator{},
	}
}

// contain information about the avatar
type Information struct {
	UserName string
}

// GenerateAndSaveAvatar generates and saves an avatar
func (s *Service) GenerateAndSaveAvatar(information Information) error {
	// defer for recovering from panic errors of the encoder
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	// encode the information
	encodeBytes, err := s.encoder.EncodeInformation(information.UserName)
	if err != nil {
		return err
	}
	// defer for recovering from panic errors of the generator and save the image
	defer func() {
		if errImage := recover(); errImage != nil {
			log.Println(errImage)
		}
	}()
	// generate and save the image
	errImage := s.generator.BuildAndSaveImage(encodeBytes)
	if errImage != nil {
		return errImage
	}
	println("The avatar has been generated")
	return nil
}
