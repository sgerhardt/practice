package trees_and_graphs

// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
//
// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
//
// Example 1:
//
// Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// Output: 6
// Explanation: The LCA of nodes 2 and 8 is 6.
// Example 2:
//
// Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
// Output: 2
// Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.
// Example 3:
//
// Input: root = [2,1], p = 2, q = 1
// Output: 2
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for (root.Val-p.Val)*(root.Val-q.Val) > 0 {
		if p.Val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

func lowestCommonAncestorRecursive(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorRecursive(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorRecursive(root.Right, p, q)
	}
	return root
}

//        if(root.val > p.val && root.val > q.val){
//            return lowestCommonAncestor(root.left, p, q);
//        }else if(root.val < p.val && root.val < q.val){
//            return lowestCommonAncestor(root.right, p, q);
//        }else{
//            return root;
//        }
//    }
