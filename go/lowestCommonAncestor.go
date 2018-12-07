/*
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (root == NULL) {
            return NULL;
        }
        if (root == p || root == q) {
            return root;
        }
        TreeNode *l = lowestCommonAncestor(root->left, p, q);
        TreeNode *r = lowestCommonAncestor(root->right, p, q);
        if (l != NULL && r != NULL && l != r) {
            return root;
        } else if (l != NULL) {
            return lowestCommonAncestor(root->left, p, q);
        } else if (r != NULL) {
            return lowestCommonAncestor(root->right, p, q);
        } else {
            return NULL;
        }
    }
    
};
*/

type TreeNode struct {
	Val int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode ,p *TreeNode, q *TreeNode) *TreeNode {
	lb := find(root.Left, p, q)
	rb := find(root.Right, p, q)
	if  lb && rb {
		return root
	}
	if lb {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if rb {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return nil
}

func find(root *TreeNode, p *TreeNode, q *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p || root == q {
		return true
	}
	return find(root.Left, p, q) || find(root.Right, p, q)
}
