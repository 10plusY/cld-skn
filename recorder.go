package cloudskine

import (
	"encoding/csv"
	"os"
	"log"
)

type Recorder struct {
	writer *csv.Writer
}

func (r *Recorder) setWriter(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			log.Println(err.Error())
		}
		r.writer = csv.NewWriter(file)
	}
}

func (r *Recorder) writeRecord(path string, note Note) {
	if r.writer == nil {
		r.setWriter(path)
	}
	defer func() {
		r.writer.Flush()
		if err := r.writer.Error(); err != nil {
			log.Println(err.Error())
		}
	}()

	r.writer.Write(note.toRecord())
}	