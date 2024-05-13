## Signature Verifier
This is a skeleton for a CLI tool that can both sign and verify signatures of files.

### Build
To build the project, run the following command:
```bash
go build -o verifier
```

### Usage
Look at the automatically generated help text with 
```bash
./verifier --help
./verifier sign --help
./verifier verify --help
```

### Task

#### Part 1: Using `openssl` CLI tool to verify signatures

1. Create a private key with the following command:
```bash
openssl genpkey -algorithm RSA -out private_key.pem -aes256
```
2. Extract the public key from the private key with the following command:
```bash
openssl rsa -pubout -in private_key.pem -out public_key.pem
```
3. Sign a file with the private key with the following command:
```bash
openssl dgst -sha256 -sign private_key.pem -out signature.sha256 secret.txt
```

> [!NOTE]
> Double-check your understanding: What is SHA256? What is it's purpose in generating the signature?

4. Verify the signature with the public key with the following command:
```bash
openssl dgst -sha256 -verify public_key.pem -signature signature.sha256 secret.txt
```
5. Modify the `secret.txt` file and verify the signature again. What happens?

#### Code Implementation
- [ ] Clone this repo and create a new branch off of `main` to work on.
- [ ] Implement the missing parts of the `verify` command. 
    - [ ] Testing - You can use the signature file + public/private key files generated in the previous step to test your implementation.
- [ ] Implement the `sign` command. The `sign` command should sign a file with a private key and save the signature to a file.
- [ ] Add a new command `generate` that generates a new private/public key pair and saves them to files. You'll need to modify the CLI/cobra commands to add this new command.

> [!IMPORTANT]
> Let me know if anything is unclear or if you need help with anything. I made this in an hour or two so there might be some mistakes or unclear instructions 😬.
> Feel free to use any resources (CLI tools, libraries, google, ChatGPT, whatever). The main goal is to familiarize yourself with Go, signing/crypto concepts, and to make sure you understand what you're doing!
> 