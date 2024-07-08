package arrays_and_string_manipulation

//Given an array nums of distinct integers,
//return all the possible permutations. You can return the answer in any order.

// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// Example 2:
//
// Input: nums = [0,1]
// Output: [[0,1],[1,0]]
// Example 3:
//
// Input: nums = [1]
// Output: [[1]]
func permute(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

//public List<List<Integer>> permute(int[] nums) {
//   List<List<Integer>> list = new ArrayList<>();
//   // Arrays.sort(nums); // not necessary
//   backtrack(list, new ArrayList<>(), nums);
//   return list;
//}
//
//private void backtrack(List<List<Integer>> list, List<Integer> tempList, int [] nums){
//   if(tempList.size() == nums.length){
//      list.add(new ArrayList<>(tempList));
//   } else{
//      for(int i = 0; i < nums.length; i++){
//         if(tempList.contains(nums[i])) continue; // element already exists, skip
//         tempList.add(nums[i]);
//         backtrack(list, tempList, nums);
//         tempList.remove(tempList.size() - 1);
//      }
//   }
//}

//
