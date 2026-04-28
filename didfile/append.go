package didfile

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Append an entry to the Didfile. Use time of last entry as timestamp.
func (d Didfile) Append(entry string) error {
	return d.append(entry, time.Time{})	
}

// Append an entry to the Didfile. Use time of last entry as timestamp. t will become time
// of last entry if it is after the current time of last entry and before the present
// moment.
func (d Didfile) AppendUntil(entry string, t time.Time) error {
	return d.append(entry, t)
}

const seekEnd int = 2

func (d Didfile) append(entry string, mtime time.Time) error {
	err := d.validate(mtime)
	if err != nil {
		return err
	}
	
	_, err = d.file.Seek(0, seekEnd)
	if err != nil {
		return err
	}
	
	_, err = fmt.Fprintf(d.file, "%v:%v\n", d.modtime.UTC().Format(time.RFC3339), entry)
	if err != nil {
		return err
	}

	if !mtime.IsZero() {
		d.modtime = mtime
		err = os.Chtimes(d.file.Name(), time.Time{}, mtime)
		if err != nil {
			return err
		}
	} else {
		d.modtime = time.Now()
	}

	return nil
}
	
func (d Didfile) validate(mtime time.Time) error {
	if mtime.IsZero() {
		return nil
	}

	if mtime.Before(d.modtime) {
		return errors.New("mtime before modtime.")
	}

	if mtime.After(time.Now()) {
		return errors.New("mtime in the future.")
	}

	return nil
}
