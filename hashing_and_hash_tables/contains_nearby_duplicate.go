package hashing_and_hash_tables

func containsNearbyDuplicate(nums []int, k int) bool {
	valuesToIndexes := map[int][]int{}
	for idx, value := range nums {
		if indexes, found := valuesToIndexes[value]; found {
			indexes = append(indexes, idx)
			valuesToIndexes[value] = indexes
		} else {
			valuesToIndexes[value] = []int{idx}
		}
	}

	for _, indexes := range valuesToIndexes {
		if len(indexes) < 2 {
			continue
		}
		// find a diff that is less than or equal to k
		for i := 0; i < len(indexes)-1; i++ {
			for j := i + 1; j < len(indexes); j++ {
				if indexes[j]-indexes[i] <= k {
					return true
				}
			}
		}
	}
	return false
}
