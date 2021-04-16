package leetcode

func dfs(root *TreeNode, pre, min *int) {
	if root == nil {
		return
	}
	dfs(root.Left, pre, min)
	if pre == nil {
		pre = &root.Val
	} else if gap := root.Val - *pre; *min > gap {
		*min = gap
	}
	dfs(root.Right, pre, min)
}

func minDiffInBST(root *TreeNode) int {
	re := int(1e5 + 1)
	dfs(root, nil, &re)
	return re
}
