package hours

import (
	"errors"
	"os"
	"time"
	"path/filepath"
)

type fileMode int

const (
	openAppend fileMode = fileMode(os.O_APPEND | os.O_WRONLY)
	openRead fileMode = fileMode(os.O_RDONLY)
)

type file struct {
	*os.File
	LastModified time.Time
}

func didfile(m fileMode) (file, error) {
	path, pathErr := path()
	if pathErr != nil {
		return file{}, pathErr
	}

	f, openErr := os.OpenFile(path, int(m), 0644)
	if openErr != nil {
		return file{}, openErr
	}

	info, infoErr := f.Stat()
	if infoErr != nil {
		return file{}, infoErr
	}

	if !info.Mode().IsRegular() {
		return file{}, errors.New("invalid didfile.")
	}

	return file{f, info.ModTime().Round(0).UTC()}, nil
}

func path() (string, error) {
	if env, envOk := os.LookupEnv("DIDFILE"); envOk {
		return env, nil
	}

	if data, dataOk := os.LookupEnv("XDG_DATA_HOME"); dataOk {
		return filepath.Join(data, "didfile"), nil
	}

	home, err := os.UserHomeDir(); 
	if err != nil {
		return "", err
	}
		
	return filepath.Join(home, "didfile"), nil
}
