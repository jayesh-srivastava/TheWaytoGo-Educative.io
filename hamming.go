package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strings differ in lengths")
	}

	dist := 0
	for x := range a {
		if a[x] != b[x] {
			dist++
		}
	}

	return dist, nil
}
