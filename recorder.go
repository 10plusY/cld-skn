package cloudskine

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Recorder struct {
	writer *csv.Writer
}

func (r *Recorder) recordRecord(path string, note Note, tagged bool, separate bool) {
	var record []string
	var file *os.File

	if tagged == true {
		record = note.ToTaggedRecord(separate)
	} else {
		record = note.ToRecord()
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err = os.Create(path)
		if err != nil {
			log.Panic(fmt.Sprintf("Error: %s", err.Error()))
		}
	}

	r.writer = csv.NewWriter(file)
	defer file.Close()
	defer r.writer.Flush()

	if err := r.writer.Write(record); err != nil {
		log.Panic(fmt.Sprintf("Error: %s", err.Error()))
	}
}
