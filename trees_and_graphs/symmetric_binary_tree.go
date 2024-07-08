package trees_and_graphs

//Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
//Input: root = [1,2,2,3,4,4,3]
//Output: true

func isSymmetric(root *BinaryNode) bool {
	return root == nil || isSymmetricHelp(root.Left, root.Right)
}

var values []int64

func isSymmetricHelp(left *BinaryNode, right *BinaryNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricHelp(left.Left, right.Right) && isSymmetricHelp(left.Right, right.Left)
}

// Java implementation. Uses a stack to keep track of tree elements. Pushes two at a time and later pops two at a time
// and checks if they are equal.
// NOTE: We push left.left, right.right then left.right and right.left to check for mirror image.
//public boolean isSymmetric(TreeNode root) {
//    if(root==null)  return true;
//
//    Stack<TreeNode> stack = new Stack<TreeNode>();
//    TreeNode left, right;
//    if(root.left!=null){
//        if(root.right==null) return false;
//        stack.push(root.left);
//        stack.push(root.right);
//    }
//    else if(root.right!=null){
//        return false;
//    }
//
//    while(!stack.empty()){
//        if(stack.size()%2!=0)   return false;
//        right = stack.pop();
//        left = stack.pop();
//        if(right.val!=left.val) return false;
//
//        if(left.left!=null){
//            if(right.right==null)   return false;
//            stack.push(left.left);
//            stack.push(right.right);
//        }
//        else if(right.right!=null){
//            return false;
//        }
//
//        if(left.right!=null){
//            if(right.left==null)   return false;
//            stack.push(left.right);
//            stack.push(right.left);
//        }
//        else if(right.left!=null){
//            return false;
//        }
//    }
//
//    return true;
//}
