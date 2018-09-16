package cloudskine

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Recorder struct {
	writer *csv.Writer
	logger *Logger
	files  []*os.File
}

func (r *Recorder) getFileFromPath(path string) *os.File {
	var rFile *os.File
	for _, file := range r.files {
		if path == file.Name() {
			rFile = file
			break
		}
	}

	return rFile
}

func (r *Recorder) getNoteRecord(note Note, tagged, separate bool) []string {
	if tagged == true {
		return note.ToTaggedRecord(separate)
	} else {
		return note.ToRecord()
	}
}

func (r *Recorder) addRecordFile(path string) {
	var file *os.File
	var err error
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err = os.Create(path)
	} else {
		file, err = os.Open(path)
	}

	if err != nil {
		log.Panic(fmt.Sprintf("Error: %s", err.Error()))
	}

	r.files = append(r.files, file)
}

func (r *Recorder) setWriterFile(path string) {
	file := r.getFileFromPath(path)
	if file == nil {
		log.Panic("File does not exist")
	}

	r.writer = csv.NewWriter(file)
}

func (r *Recorder) writeNoteRecord(path string, record []string) {
	file := r.getFileFromPath(path)
	if file == nil {
		log.Panic("File does not exist")
	}

	r.writer = csv.NewWriter(file)
	defer file.Close()
	defer r.writer.Flush()

	err := r.writer.Write(record)
	if err != nil {
		log.Panic("Cannot write record")
	}
}
