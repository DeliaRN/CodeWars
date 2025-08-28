package kata04hours

import "fmt"

func HumanReadableTime(seconds int) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60

	result := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, secs)
	return result
}
