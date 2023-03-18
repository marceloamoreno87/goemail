package local

import (
	"log"
	"mime/multipart"
	"os"
)

func Delete(attachment *multipart.FileHeader) {
	filepath := "/tmp/attachments/" + attachment.Filename
	e := os.Remove(filepath)
	if e != nil {
		log.Fatal(e)
	}
}
