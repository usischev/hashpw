package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	saltBytes        int = 16
	hashBytes        int = 64
	pbkdf2Iterations int = 100000
)

var b64Replacer = strings.NewReplacer("+", ".")

func createHash(password string) (string, error) {
	salt := make([]byte, saltBytes)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := pbkdf2.Key([]byte(password), salt, pbkdf2Iterations, hashBytes, sha512.New)
	saltString := b64Replacer.Replace(base64.RawStdEncoding.EncodeToString(salt))
	hashString := b64Replacer.Replace(base64.RawStdEncoding.EncodeToString(hash))

	return fmt.Sprintf("{PBKDF2-SHA512}%d$%s$%s", pbkdf2Iterations, saltString, hashString), err
}

func main() {
	var passwd string

	// 	read directly from stdin if pipe, prompt and hide user input otherwise
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		reader := bufio.NewReader(os.Stdin)
		passwd, _ = reader.ReadString('\n')
		passwd = strings.Replace(passwd, "\n", "", -1)
	} else {
		fmt.Print("Password:")
		bytePasswd, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			panic(err)
		}
		fmt.Println("")
		passwd = string(bytePasswd)
	}

	pwHash, _ := createHash(passwd)
	fmt.Println(pwHash)
}
