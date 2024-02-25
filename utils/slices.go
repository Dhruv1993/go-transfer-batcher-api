package utils

// passing a map of stringed key and value pair,it will return an array of keys
func MapKeys(m map[string]string) []string {
	i := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}