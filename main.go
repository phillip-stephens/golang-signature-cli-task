package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "./verifier --help",
	Short: "a file signer and verifier",
}

var fileName, publicKeyFileName, privateKeyFileName, signatureFileName string

// add 2 subcommands to the root command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "sign a file",
	Long:  "sign a file using a private key",
	Run: func(cmd *cobra.Command, args []string) {
		sign(fileName, privateKeyFileName, signatureFileName)
	},
}

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify a file",
	Long:  "verify a file using a public key",
	Run: func(cmd *cobra.Command, args []string) {
		verify(fileName, publicKeyFileName, signatureFileName)
	},
}

func main() {
	signCmd.Flags().StringVarP(&fileName, "file", "f", "", "file to sign")
	signCmd.Flags().StringVarP(&privateKeyFileName, "private-key", "k", "", "private key file")
	signCmd.Flags().StringVarP(&signatureFileName, "signature", "o", "", "signature file")

	verifyCmd.Flags().StringVarP(&fileName, "file", "f", "", "file to verify")
	verifyCmd.Flags().StringVarP(&publicKeyFileName, "public-key", "k", "", "public key file")
	verifyCmd.Flags().StringVarP(&signatureFileName, "signature", "s", "", "signature file")

	rootCmd.AddCommand(signCmd)
	rootCmd.AddCommand(verifyCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
