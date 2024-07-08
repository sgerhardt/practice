package trees_and_graphs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var queue []*TreeNode
	queue = append(queue, root)

	var result [][]int
	for len(queue) > 0 {
		var level []int
		cnt := len(queue)
		// this inner loop is the key to adding the level results correctly... it allows for fanning out to the needed number of nodes
		for i := 0; i < cnt; i++ {
			node := queue[0]
			if len(queue) == 1 {
				queue = []*TreeNode{}
			} else {
				queue = queue[1:]
			}

			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}

//List<List<Integer>> res = new ArrayList<>();
//if (root == null) return res;
//Queue<TreeNode> queue = new LinkedList<>();
//queue.add(root);
//while (!queue.isEmpty()) {
//　　List<Integer> level = new ArrayList<>();
//　　int cnt = queue.size();
//　　for (int i = 0; i < cnt; i++) {
//　　　　TreeNode node = queue.poll();
//　　　　level.add(node.val);
//　　　　if (node.left != null) {
//　　　　　　queue.add(node.left);
//　　　　}
//　　　　if (node.right != null) {
//　　　　　　queue.add(node.right);
//　　　　}
//　　}
//　　res.add(level);
//}
//return res;
