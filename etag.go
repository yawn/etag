package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yawn/etag/hash"
)

var size uint64

func init() {

	flag.Uint64Var(&size, "chunksize", 8*1024*1024, "size of the multipart chunks")

	flag.Parse()

}

func main() {

	for _, path := range flag.Args() {

		tag, err := hash.New(path, size)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("%s\t%s\n", path, *tag)

	}

}
