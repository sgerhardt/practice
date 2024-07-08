package arrays_and_string_manipulation

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
// A subarray is a contiguous part of an array.
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Apparently, this is a optimization problem, which can be usually solved by DP. So when it comes to DP, the first thing for us to figure out is the format of the sub problem(or the state of each sub problem). The format of the sub problem can be helpful when we are trying to come up with the recursive relation.
//
// At first, I think the sub-problem should look like: maxSubArray(int A[], int i, int j), which means the maxSubArray for A[i: j]. In this way, our goal is to figure out what maxSubArray(A, 0, A.length - 1) is. However, if we define the format of the sub problem in this way, it's hard to find the connection from the sub problem to the original problem(at least for me). In other words, I can't find a way to divided the original problem into the sub problems and use the solutions of the sub problems to somehow create the solution of the original one.
//
// So I change the format of the sub-problem into something like: maxSubArray(int A[], int i), which means the maxSubArray for A[0:i ] which must has A[i] as the end element.
// Note that now the sub problem's format is less flexible and less powerful than the previous one because there's a limitation that A[i] should be contained in that sequence and we have
// to keep track of each solution of the sub problem to update the global optimal value. However, now the connection between the sub problem & the original one becomes clearer
func maxSubArray(nums []int) int {

	// We store values relevant to calculating the max sub array in the dp (dynamic programming) slice.
	// It will be either the previous value plus the current value if that value is greater than zero,
	// or just use the existing value of nums[i].
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	maxVal := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		maxVal = max(maxVal, dp[i])
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
