package encoder

import (
	"crypto/md5"
	"log"
)

// interface para encodear información
type CryptoEncoder struct{}

// Metodo para encodear información de CryptoEcoder.
// Recibe un string y devuelve un array de bytes con la información encodeada en md5.
func (e *CryptoEncoder) EncodeInformation(strInformation string) (EncodeInformation []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	hash := md5.New()
	hash.Write([]byte(strInformation))
	EncodeInformation = hash.Sum(nil)
	return EncodeInformation, nil
}
