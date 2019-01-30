package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/ssh"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <public.asc>\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, ">>> opening source file %s\n", os.Args[1])
	sourceFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not open source file: %s\n", err.Error())
		os.Exit(2)
	}
	defer sourceFile.Close()


	fmt.Fprintln(os.Stderr, ">>> decoding OpenPGP keyring")
	entityList, err := openpgp.ReadArmoredKeyRing(sourceFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not parse source file: %s\n", err.Error())
		os.Exit(3)
	}

	if len(entityList) != 1 {
		fmt.Fprintf(os.Stderr, "ERROR: expected exactly one key in source file, got %d\n", len(entityList))
		os.Exit(4)
	}

	key := entityList[0]
	fmt.Fprintf(os.Stderr, ">>> converting key %s\n", key.PrimaryKey.KeyIdString())

	sshPubKey, err := ssh.NewPublicKey(key.PrimaryKey.PublicKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not convert to SSH public key: %s\n", err.Error())
		os.Exit(5)
	}

	fmt.Fprintln(os.Stderr, ">>> converted key (on stdout): ")
	os.Stdout.Write(ssh.MarshalAuthorizedKey(sshPubKey))
}
