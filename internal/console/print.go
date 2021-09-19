package console

import (
	"log"

	"github.com/jrbeverly/golang-analyzer-inline-bazel/internal/cobrago"
)

type ConsoleWriter struct {
}

func NewConsoleWriter() ConsoleWriter {
	writer := ConsoleWriter{}
	return writer
}

func (r ConsoleWriter) Print(files []cobrago.RemoteFile) {
	log.Println("List of files:")
	for _, object := range files {
		log.Printf("key=%s size=%d\n", object.Key, object.Size)
	}
}
