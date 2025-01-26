// Package record is responsible for recording the calculated investment score and its time of calculation.
package record

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

const FILE_NAME = "investment_score_record.csv"

const TIME_FORMAT = "2006-01-02"

// TODO: Implement io.Writer if this fucntion must be named as Write
func Write(score int) error {
	now := time.Now()

	var f *os.File
	var newFileCreated bool

	if _, err := os.Stat(FILE_NAME); os.IsNotExist(err) {
		f, err = os.Create(FILE_NAME)
		if err != nil {
			return fmt.Errorf("failed to create new investment score record file: %w", err)
		}
		newFileCreated = true
	} else {
		f, err = os.OpenFile(FILE_NAME, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return fmt.Errorf("failed to open an existing insvestment score record file: %w", err)
		}
	}

	w := csv.NewWriter(f)

	// Add header to new file
	if newFileCreated {
		_ = w.Write([]string{"date", "investment score"})
	}
	_ = w.Write([]string{now.Format(TIME_FORMAT), strconv.Itoa(score)})
	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}
