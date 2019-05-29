package template

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// There must be two writers: FULL and APPEND
// Each accepting a rendered template
type Writer struct {
	Path string
	data []byte
}

type FileWriter struct {
	Writer
}

// NewFileWriter initializes a new FileWriter
func NewFileWriter(path string, data []byte) *FileWriter {
	return &FileWriter{
		Writer{
			Path:path,
			data:data,
		},
	}
}

// Write dumps the given data into a file and creates it if necessary.
// The overwrite flag can be set to overwrite any existing data.
func (f *FileWriter) Write(overwrite bool) error {
	if _, err := os.Stat(f.Path); err == nil {
		if !overwrite {
			return fmt.Errorf("%s file already exists, overwrite is disabled", f.Path)
		}
	}

	if err := ioutil.WriteFile(f.Path, f.data, 0644); err != nil {
		return errors.Wrap(err, "failed to write file")
	}
	return nil
}

type FileAppendWriter struct {
	Writer
}

// NewFileAppendWriter returns a new appending file-writer for Godin templates
func NewFileAppendWriter(path string, data []byte) *FileAppendWriter {
    return &FileAppendWriter{
    	Writer{
    		data:data,
    		Path:path,
		},
	}
}

// Write will open the given file and try to append the given data to it
// The file is NOT created if it doesn't exist.
func (f *FileAppendWriter) Write() error {
	file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
	    return errors.Wrap(err, "file to append cannot be opened")
	}

	if _, err := file.Write(f.data); err != nil {
		return err
	}
	return nil	
}