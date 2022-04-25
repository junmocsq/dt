package part4

import (
	"errors"
	"fmt"
)

type BinTreeNode struct {
	Elem  interface{}
	Left  *BinTreeNode
	Right *BinTreeNode
}

type BinTree struct {
	Root    *BinTreeNode
	NodeNum int
}

type BinaryTree interface {
	MakeEmpty() bool
	Find(interface{}) *BinTreeNode
	FindMin(*BinTreeNode) (*BinTreeNode, error)
	FindMax(*BinTreeNode) (*BinTreeNode, error)
	Insert(interface{}) error
	Delete(interface{}) error

	Compare(interface{}, interface{}) (int, error)
}

func NewTree() *BinTree {
	return &BinTree{}
}

func (t *BinTree) Compare(a, b interface{}) (int, error) {
	return Compare(a, b)
}

func (t *BinTree) MakeEmpty() {
	t.Root = nil
	t.NodeNum = 0
}

func (t *BinTree) Find(ele interface{}) (*BinTreeNode, error) {
	if t.Root == nil {
		return nil, errors.New("搜索树为空")
	}
	return t.find(ele, t.Root)
}

func (t *BinTree) find(ele interface{}, node *BinTreeNode) (*BinTreeNode, error) {
	res, err := t.Compare(ele, node.Elem)
	if err != nil {
		return nil, err
	}
	if res == 1 {
		if node.Right == nil {
			return nil, errors.New("搜索元素不存在")
		}
		return t.find(ele, node.Right)
	} else if res == -1 {
		if node.Left == nil {
			return nil, errors.New("搜索元素不存在")
		}
		return t.find(ele, node.Left)
	} else {
		return node, nil
	}
}

func (t *BinTree) FindMin(node *BinTreeNode) (*BinTreeNode, error) {
	if node == nil {
		return nil, errors.New("搜索树为空")
	}

	for node.Left != nil {
		node = node.Left
	}
	return node, nil
}

func (t *BinTree) FindMax(node *BinTreeNode) (*BinTreeNode, error) {
	if node == nil {
		return nil, errors.New("搜索树为空")
	}

	for node.Right != nil {
		node = node.Right
	}
	return node, nil
}

func (t *BinTree) Insert(ele interface{}) error {
	node := &BinTreeNode{Elem: ele}
	if t.Root == nil {
		t.Root = node
		t.NodeNum = 1
		return nil
	}
	root := t.Root
	err := t.insert(ele, root)
	return err
}

func (t *BinTree) insert(ele interface{}, node *BinTreeNode) error {
	res, err := t.Compare(ele, node.Elem)
	if err != nil {
		return err
	}
	if res == 1 {
		if node.Right == nil {
			node.Right = &BinTreeNode{Elem: ele}
			t.NodeNum++
			return nil
		}
		return t.insert(ele, node.Right)
	} else if res == -1 {
		if node.Left == nil {
			node.Left = &BinTreeNode{Elem: ele}
			t.NodeNum++
			return nil
		}
		return t.insert(ele, node.Left)
	} else {
		return nil
	}
}

func (t *BinTree) Delete(ele interface{}) error {
	if t.Root == nil {
		return errors.New("树为空")
	}
	err := t.delete(ele, t.Root, t.Root)
	if err != nil {
		return err
	}
	t.NodeNum--
	return nil
}
func (t *BinTree) delete(ele interface{}, node, parent *BinTreeNode) error {
	res, err := t.Compare(ele, node.Elem)
	if err != nil {
		return err
	}
	if res == 1 {
		if node.Right == nil {
			return errors.New("不存在要删除的元素")
		}
		return t.delete(ele, node.Right, node)
	} else if res == -1 {
		if node.Left == nil {
			return errors.New("不存在要删除的元素")
		}
		return t.delete(ele, node.Left, node)
	} else {
		// 相等
		if node.Left != nil && node.Right != nil {
			temp, err := t.FindMin(node.Right)
			if err != nil {
				return err
			}
			node.Elem = temp.Elem
			return t.delete(temp.Elem, node.Right, node)
		} else if node.Left == nil && node.Right == nil {
			if node == parent {
				// 根节点
				t.Root = nil
			} else {
				if parent.Right == node {
					parent.Right = nil
				} else {
					parent.Left = nil
				}
			}
		} else {
			if node.Right == nil {
				*node = *node.Left
			} else {
				*node = *node.Right
			}
		}
	}
	return nil
}

func (t *BinTree) PrintTree() {
	fmt.Println("tree: 节点个数:", t.NodeNum)
	t.printtree(t.Root)
}

func (t *BinTree) printtree(node *BinTreeNode) {
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
