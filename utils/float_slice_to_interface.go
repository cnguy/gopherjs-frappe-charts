package utils

// FloatSliceToInterface helps convert a float64 array to an interface{}.
func FloatSliceToInterface(values []float64) []interface{} {
	var newLabels []interface{}
	newLabels = make([]interface{}, len(values))
	for i, v := range values {
		newLabels[i] = v
	}
	return newLabels
}
