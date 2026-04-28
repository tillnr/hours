package didfile

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"
)

// Where all tracked time is stored.
type Didfile struct {
	file *os.File
	modtime time.Time
}

// Choose the first of the following locations and open the file located there:
// - $DIDFILE
// - $XDG_DATA_HOME/didfile
// - os.UserHomeDir()/didfile
// Return error when no location can be found, opening the file at location fails or the
// opened file is not a regular file.
func Open() (Didfile, error) {
	path, pathErr := path()
	if pathErr != nil {
		return Didfile{}, pathErr
	}

	f, openErr := os.OpenFile(path, os.O_RDWR, 0644)
	if openErr != nil {
		return Didfile{}, openErr
	}

	info, infoErr := f.Stat()
	if infoErr != nil {
		return Didfile{}, infoErr
	}

	if !info.Mode().IsRegular() {
		return Didfile{}, errors.New("invalid didfile.")
	}

	return Didfile{f, info.ModTime()}, nil
}

func (d Didfile) ModTime() time.Time {
	return d.modtime
}

func (d Didfile) Reader() io.Reader {
	return d.file
}

func (d Didfile) ReadSeekWriter() io.ReadWriteSeeker {
	return d.file
}

func (d Didfile) Close() {
	d.file.Close()
}

func path() (string, error) {
	if env, envOk := os.LookupEnv("DIDFILE"); envOk {
		return env, nil
	}

	if data, dataOk := os.LookupEnv("XDG_DATA_HOME"); dataOk {
		return filepath.Join(data, "didfile"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, "didfile"), nil
}
