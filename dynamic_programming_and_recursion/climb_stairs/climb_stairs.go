package climb_stairs

// You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Example 1:
//
// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
// Example 2:
//
// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step
func climbStairs(n int) int {
	currentStep := 0
	previousStep := 1
	totalSteps := 0
	for i := n; i > 0; i-- {
		totalSteps = currentStep + previousStep
		currentStep = previousStep
		previousStep = totalSteps
	}
	return totalSteps
}

//We take a bottom-up dynamic programming approach. Say we are on step 5 of 5. Then there are 0 steps to go.
// If we are on step 4, there is one step to go. Step 3 is the result of the last answer (1) plus the total steps so far (1).
// So step three has two ways to climb the remaining steps. Step 2 likewise has three ways to go. Step 1 likewise has 5 ways to go.
// And finally step 0 has 8 ways to go (results from step 1 and 2)
