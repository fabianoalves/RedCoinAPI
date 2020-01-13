package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const chaveCripto string = "redcoinApi2020!@"

//CriptografarSenha retorna o valor criptografado para senhas usadas na API
func CriptografarSenha(senha string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(senha), 14)
	return string(bytes)
}

//VerificarSenha retorna a comparação entre a senha e o has criptografado
func VerificarSenha(senha, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	return err == nil
}

//DecriptografarChaves retorna o valor original do texto criptografado
func DecriptografarChaves(chave string) string {
	textoCodificado, _ := base64.URLEncoding.DecodeString(chave)
	block, _ := aes.NewCipher([]byte(chaveCripto))

	iv := textoCodificado[:aes.BlockSize]
	textoCodificado = textoCodificado[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(textoCodificado, textoCodificado)

	return fmt.Sprintf("%s", textoCodificado)
}
