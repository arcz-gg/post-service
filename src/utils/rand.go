package utils;

import (
	"crypto/rand"
	"io"
	"log"

)

func GenId(length int) string {

	b := make([]byte, length);
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'};

	nbrs, err := io.ReadAtLeast(rand.Reader, b, length);
	if nbrs != length {
		log.Fatal(err);
	}

	for i, vals := range b {
		
		b[i] = table[int(vals)%len(table)];
	}

	return string(b);
}
