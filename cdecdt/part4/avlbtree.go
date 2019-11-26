package part4

import "fmt"

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
	if this.Root == nil {
		this.Root = &AvlNode{Elem: ele}
		this.Nodenum++
	} else {
		this.insert(ele, this.Root, nil, nil)
	}
}

func (this *Avl) insert(ele int, node, parent, grandfather *AvlNode) *AvlNode {

	if node == nil {
		return &AvlNode{Elem: ele}
	} else if ele == node.Elem {
		return nil
	} else if ele < node.Elem {
		node.Left = this.insert(ele, node.Left, node, parent)
		if parent != nil && this.height(parent.Left)-this.height(parent.Right) == 2 {
			if parent.Left.Elem > ele {
				if grandfather == nil {
					this.Root = this.rightSingleRotation(parent)
				} else if grandfather.Left == parent {
					grandfather.Left = this.rightSingleRotation(parent)
				} else {
					grandfather.Right = this.rightSingleRotation(parent)
				}
			} else {
				if grandfather == nil {
					this.Root = this.leftDoubleRotation(parent)
				} else if grandfather.Left == parent {
					grandfather.Left = this.leftDoubleRotation(parent)
				} else {
					grandfather.Right = this.leftDoubleRotation(parent)
				}
			}
		}

	} else {
		node.Right = this.insert(ele, node.Right, node, parent)
		if parent != nil && this.height(parent.Right)-this.height(parent.Left) == 2 {
			if parent.Right.Elem < ele {
				if grandfather == nil {
					this.Root = this.leftSingleRotation(parent)
				} else if grandfather.Left == parent {
					grandfather.Left = this.leftSingleRotation(parent)
				} else {
					grandfather.Right = this.leftSingleRotation(parent)
				}
			} else {
				if grandfather == nil {
					this.Root = this.rightDoubleRotation(parent)
				} else if grandfather.Left == parent {
					grandfather.Left = this.rightDoubleRotation(parent)
				} else {
					grandfather.Right = this.rightDoubleRotation(parent)
				}
			}
		}
	}
	node.Height = this.max(this.height(node.Right), this.height(node.Left)) + 1
	return node
}

// 右旋
func (this *Avl) rightSingleRotation(k2 *AvlNode) *AvlNode {
	k1 := k2.Right
	k2.Right = k1.Left
	k1.Left = k2
	k2.Height = this.max(this.height(k2.Right), this.height(k2.Left)) + 1
	k1.Height = this.max(this.height(k1.Left), this.height(k1.Right)) + 1
	return k1
}

// 左旋
func (this *Avl) leftSingleRotation(k2 *AvlNode) *AvlNode {
	k1 := k2.Left
	k2.Left = k1.Right
	k1.Right = k2
	k2.Height = this.max(this.height(k2.Right), this.height(k2.Left)) + 1
	k1.Height = this.max(this.height(k1.Left), this.height(k1.Right)) + 1
	return k1
}

func (this *Avl) max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (this *Avl) height(node *AvlNode) int {
	if node == nil {
		return -1
	} else {
		return node.Height
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

func (t *Avl) PrintTree() {
	fmt.Println("tree: 节点个数:", t.Nodenum)
	t.printtree(t.Root)
}

func (t *Avl) printtree(node *AvlNode) {
	if node != nil {
		fmt.Printf("%v\t", node.Elem)
		if node.Left != nil {
			t.printtree(node.Left)
		}
		if node.Right != nil {
			t.printtree(node.Right)
		}
	}
}
