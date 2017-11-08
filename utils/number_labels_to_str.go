package utils

import "strconv"

// NumberLabelsToStr helps convert a number array (ex: of years) to string labels.
func NumberLabelsToStr(labels []int) []string {
	newLabels := make([]string, len(labels))
	for i := 0; i < len(labels); i++ {
		newLabels[i] = strconv.Itoa(labels[i])
	}
	return newLabels
}
