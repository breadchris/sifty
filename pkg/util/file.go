package util

import (
	"fmt"
	"time"
)

// GenerateFilename generates a filename with the current time.
func GenerateFilename(basename, ext string) string {
	// Get the current time.
	now := time.Now()

	// Format the current time as a string.
	timeStr := now.Format("2006-01-02_15-04-05")

	// Combine the time string with a filename extension.
	filename := fmt.Sprintf("%s-%s.%s", basename, timeStr, ext)

	return filename
}
