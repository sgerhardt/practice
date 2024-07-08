package hashing_and_hash_tables

//Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it.
// That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].
//
//Return the answer in an array.
//Input: nums = [8,1,2,2,3]
//Output: [4,0,1,1,3]
//Explanation:
//For nums[0]=8 there exist four smaller numbers than it (1, 2, 2 and 3).
//For nums[1]=1 does not exist any smaller number than it.
//For nums[2]=2 there exist one smaller number than it (1).
//For nums[3]=2 there exist one smaller number than it (1).
//For nums[4]=3 there exist three smaller numbers than it (1, 2 and 2).

func smallerNumbersThanCurrent(nums []int) []int {
	count := make(map[int]int, 101)
	res := make([]int, len(nums))

	// first loop to track occurrences of each number
	for _, n := range nums {
		count[n]++
	}

	// We've counted our hits so far, now go through them from 1 to 100 and do a running count.
	// Any later numbers are naturally larger than earlier numbers so the _running_ count works here...
	//e.g.  count[0]==0,  count[1]==1 + 0, count[2]==2 + 1 + 0, count[3] == 1 + 2 + 1 + 0 and so on...
	for i := 1; i <= 100; i++ {
		count[i] += count[i-1]
	}

	// we know how many occurrences of each number exist at each index from 0 to 100,
	// so the answer to how many are smaller than n is simply the count at the previous number, count[nums[i]-1]
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			res[i] = 0
		} else {
			res[i] = count[nums[i]-1]
		}
	}
	return res
}
