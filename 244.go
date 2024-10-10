// Package dog helps convert years
package dog

// Returns the converted value when converting human years to dog years
func convertToDog(i float64) float64 {
	v := i * 7
	return v
}

// Returns the converted value when converting dog years to human years
func convertToHuman(i float64) float64 {
	v := i / 7
	return v
}