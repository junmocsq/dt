package part4

// 一棵AVL树是其每个节点的左子树和右子树的高度最多差1的二叉查找树。

type AvlNode struct {
	Elem   int
	Left   *AvlNode
	Right  *AvlNode
	Height int
}

type Avl struct {
	Nodenum int
	Root    *AvlNode
}

func (this *Avl) Insert(ele int) {

}

// 右旋
func (this *Avl) rightSingleRotation(k2 *AvlNode) *AvlNode {
	k1 := k2.Right
	k2.Right = k1.Left
	k1.Left = k2
	k2.Height = max(k2.Right.Height, k2.Left.Height)
	k1.Height = max(k1.Left.Height, k1.Right.Height)
	return k1
}

// 左旋
func (this *Avl) leftSingleRotation(k2 *AvlNode) *AvlNode {
	k1 := k2.Left
	k2.Left = k1.Right
	k1.Right = k2
	k2.Height = max(k2.Right.Height, k2.Left.Height)
	k1.Height = max(k1.Left.Height, k1.Right.Height)
	return k1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 双旋 左到右
func (this *Avl) leftDoubleRotation(k3 *AvlNode) *AvlNode {
	// 左旋
	k3.Left = this.leftSingleRotation(k3.Left)
	// 右旋
	return this.rightSingleRotation(k3)
}

// 双旋 右到左
func (this *Avl) rightDoubleRotation(k3 *AvlNode) *AvlNode {
	// 右旋
	k3.Right = this.rightSingleRotation(k3.Right)
	// 左旋
	return this.leftSingleRotation(k3)
}
