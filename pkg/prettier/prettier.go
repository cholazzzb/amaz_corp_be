package prettier

import (
	"fmt"
	"strings"
)

// PrintSlice prints a slice with a custom format for any type.
func PrintSlice[T any](slice []T) {
	fmt.Println(formatSlice(slice))
}

// Helper function to format slice elements
func formatSlice[T any](slice []T) string {
	var sb strings.Builder
	sb.WriteString("[ ")

	for i, elem := range slice {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", elem))
	}

	sb.WriteString(" ]")
	return sb.String()
}
