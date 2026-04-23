package hours

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Add an entry to didfile that activity was started the last time didfile was modified.
func Add(activity string) error {
	trimmed := strings.TrimSpace(activity)
	if len(trimmed) == 0 {
		return errors.New("entry empty.")
	}

	file, openErr := didfile(openAppend)
	if openErr != nil {
		return openErr
	}

	defer file.Close()
	entry := fmt.Sprintf("%v:%v\n", file.LastModified.Format(time.RFC3339), activity)
	if _, writeErr := file.Write([]byte(entry)); writeErr != nil {
		return writeErr
	}

	return nil
}
