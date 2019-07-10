package hash

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// New creates an S3 etag for the given file and the given multipart chunksize
func New(path string, chunksize uint64) (*string, error) {

	fh, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer fh.Close()

	var (
		buf    = make([]byte, chunksize)
		chunks = 0
		etag   = md5.New()
		last   []byte
		result string
	)

	for {

		n, err := fh.Read(buf)

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		h := md5.New()
		h.Write(buf[:n])

		chunks++

		last = h.Sum(nil)
		etag.Write(last)

	}

	if chunks > 1 {
		result = fmt.Sprintf("%x-%d", etag.Sum(nil), chunks)
	} else {
		result = fmt.Sprintf("%x", last)
	}

	return &result, nil

}
