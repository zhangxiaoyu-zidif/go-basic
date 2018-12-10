

```go
// search two arrays and return true if they have at least one common element; return false otherwise
func haveOverlap(a1, a2 []string) bool {
	m := map[string]bool{}

	for _, val := range a1 {
		m[val] = true
	}
	for _, val := range a2 {
		if _, ok := m[val]; ok {
			return true
		}
	}

	return false
}
```
