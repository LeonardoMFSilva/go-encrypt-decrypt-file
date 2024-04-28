package main

import (
	"bytes"
	"fmt"
	filecrypt "github.com/akhilsharma90/go-file-encryption/fileencrypt"
	"os"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]
	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run encrypt to encrypt a file, and decrypt to decrypt a file!")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("file encryption")
	fmt.Println("simple file encryption for your day-to-dau needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tgo run . encrypt /path/to/your/file ")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypt a file given a password.")
	fmt.Println("\t decrypt\tDecrypt a file given a password.")
	fmt.Println("\t help\tDisplays help text.")
	fmt.Println("")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		println("missing the path to the file. More info at, run go run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File is not valid/found")
	}
	password := getPassword()
	fmt.Println("\nEncrypting...")
	filecrypt.EncryptFile(file, password)
	fmt.Println("\nEncrypted successfully!")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		println("missing the path to the file. More info at, run go run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File is not valid/found")
	}
	fmt.Print("Enter password:")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nDecrypting...")
	filecrypt.DecryptFile(file, password)
	fmt.Println("\nDecrypted successfully!")
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func getPassword() []byte {
	fmt.Println("Enter password: ")
	password, _ := term.ReadPassword(0)
	fmt.Print("\nConfirm password: ")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("\nInvalid password: ")
		return getPassword()
	}
	return password
}

func validatePassword(password []byte, confirmPassword []byte) bool {
	if !bytes.Equal(password, confirmPassword) {
		return false
	}
	return true
}
