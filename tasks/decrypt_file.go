package tasks

/*
   Copyright 2018 TheRedSpy15

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TheRedSpy15/Multi-Go/utils"
	"github.com/daviddengcn/go-colortext"
)

// DecryptFile decrypts the target file
// BUG: decrypted file is unusable
// NOTE: decrypt file doesn't actually save as unencrypted
func DecryptFile(target string) {
	utils.CheckTarget(target)       // make sure target is valid
	ct.Foreground(ct.Yellow, false) // set text color to dark yellow

	print("Enter Password: ")
	password := utils.GetPassword() // get password securely

	file, err := os.Create(target) // create file object
	if err != nil {
		ct.Foreground(ct.Red, true) // set text color to bright red
		panic(err.Error())
	}
	defer file.Close()                        // makes sure file gets closed
	file.Write(decryptFile(target, password)) // decrypt file
	fmt.Println("\nFile decrypted!")
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		ct.Foreground(ct.Red, true)
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		ct.Foreground(ct.Red, true)
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		ct.Foreground(ct.Red, true)
		panic(err.Error())
	}
	return plaintext
}

func decryptFile(filename string, passphrase string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		ct.Foreground(ct.Red, true)
		panic(err.Error())
	}

	return decrypt(data, passphrase)
}
