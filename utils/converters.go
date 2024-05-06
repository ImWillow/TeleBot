package utils

import "strconv"

func StringsToInts(str []string) ([]int64, error) {
	ints := []int64{}
	for _, s := range str {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		ints = append(ints, int64(i))
	}

	return ints, nil
}
