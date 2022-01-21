package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func decrypt(cipherstring string, keystring string) string {
	ciphertext := []byte(cipherstring)

	key := []byte(keystring)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("Text is too short")
	}

	iv := ciphertext[:aes.BlockSize]

	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	writeToFile(string(ciphertext), "arquivo.txt")

	return string(ciphertext)
}

func encrypt(keystring string, filename string) string {
	data, _ := ioutil.ReadFile(filename)

	plaintext := []byte(data)

	key := []byte(keystring)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return string(ciphertext)
}

func readline() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(line)
}

func writeToFile(data, file string) {
	ioutil.WriteFile(file, []byte(data), 777)
}

func readFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}

func main() {
	key := "1234567812345678"
	filename := "arquivo.txt"

	for {
		fmt.Println("Escolha uma opcao")
		fmt.Println("1 - Carregar arquivo")
		fmt.Println("2 - Criptografar")
		fmt.Println("3 - Descriptografar")
		fmt.Println("4 - Mostrar arquivo")
		fmt.Println("5 - Sair")
		line := readline()

		switch line {
		case "1":
			fmt.Println("Arquivo Carregado...")
		case "2":
			ciphertext := encrypt(key, filename)
			writeToFile(ciphertext, filename)
			fmt.Println(ciphertext)
		case "3":
			ciphertext, _ := readFromFile(filename)
			plaintext := decrypt(string(ciphertext), key)
			fmt.Println(plaintext)
		case "4":
			ciphertext, _ := readFromFile(filename)
			fmt.Println(string(ciphertext))
		case "5":
			os.Exit(0)
		}
	}
}
