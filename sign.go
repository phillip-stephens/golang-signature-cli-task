package main

import (
	"fmt"
	"os"
)

// sign signs a file using a private key
// fileNameToSign is the name of the file to sign using the private key
// privateKeyFileName is the name of the private key file used for signing
// signatureFileName is the name of the signature file that the signature will be written to
func sign(fileNameToSign, privateKeyFileName, signatureFileName string) {
	if len(fileName) == 0 {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}
	if len(privateKeyFileName) == 0 {
		fmt.Println("Please provide a private key file name")
		os.Exit(1)
	}
	fmt.Println("Sign function called against file: ", fileName)
	fmt.Println("Private key file: ", privateKeyFileName)
	// Add your code here to generate a signature for the file and print it out
	// Steps
	// 1. Read the file
	// 2. Generate a hash of the file
	// 3. Read the private key
	// 4. Sign the hash using the private key
	// 5. Write the signature to a file
}

// verify verifies a signature for a file matches a hash of the file contents using a public key
func verify(fileNameToVerify, publicKeyFileName, signatureFileName string) {
	if len(fileNameToVerify) == 0 {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}
	if len(publicKeyFileName) == 0 {
		fmt.Println("Please provide a public key file name")
		os.Exit(1)
	}
	if len(signatureFileName) == 0 {
		fmt.Println("Please provide a signature file name")
		os.Exit(1)
	}
	fmt.Println("Verify function called against file: ", fileNameToVerify)
	fmt.Println("Public key file: ", publicKeyFileName)
	fmt.Println("Signature file: ", signatureFileName)
	// Add your code here to verify the signature for the file and print out the result
	// Steps
	// 1. Read the file
	// 2. Generate a hash of the file
	// 3. Read the public key
	// 4. Read the signature
	// 5. Verify the signature using the public key
}
