package google

import (
	"log"
	"testing"
)

func TestGoogleAuth_GetSecret(t *testing.T) {
	auth := NewGoogleAuth()
	log.Println(auth.GetSecret())

}

func TestGoogleAuth_GetCode(t *testing.T) {

	auth := NewGoogleAuth()
	//secret := auth.GetSecret()
	code, _ := auth.GetCode("A4CICLQRAOY63I7GUOZJRJVUR6E74EAO")

	log.Println(code)
}

func TestGoogleAuth_GetQrString(t *testing.T) {
	auth := NewGoogleAuth()
	secret := auth.GetSecret()
	log.Println(secret)
	log.Println(auth.GetQrString(secret))
}
