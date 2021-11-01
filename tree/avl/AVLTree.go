package avl

import "fmt"

type AVLTreeNode struct {
	Data      int
	Parent    *AVLTreeNode
	LeftTree  *AVLTreeNode
	RightTree *AVLTreeNode
}

type AVLTree struct {
	Root *AVLTreeNode
}

func BuildTree() *AVLTree {
	return &AVLTree{nil}
}

func BuildTreeByData(data int) *AVLTree {
	return &AVLTree{buildNode(data)}
}

func buildNode(data int) *AVLTreeNode {
	return &AVLTreeNode{data, nil, nil, nil}
}

func buildNodeByDataAndParent(data int, parent *AVLTreeNode) *AVLTreeNode {
	return &AVLTreeNode{data, parent, nil, nil}
}

func getTreeDepth(node *AVLTreeNode) int {
	if node == nil {
		return 0
	}
	leftTreeDepth := getTreeDepth(node.LeftTree)
	rightTreeDepth := getTreeDepth(node.RightTree)
	return getMax(leftTreeDepth, rightTreeDepth) + 1
}

func getMax(num1, num2 int) (max int) {
	if num1 >= num2 {
		max = num1
		return
	}
	max = num2
	return
}

// 计算平衡因子
func calcBalanceFactor(node *AVLTreeNode) int {
	if node == nil {
		return 0
	}
	return getTreeDepth(node.LeftTree) - getTreeDepth(node.RightTree)
}

// 左旋
func (t *AVLTree) leftRotation(node *AVLTreeNode) {
	if node == nil {
		return
	}
	rightChild := node.RightTree
	node.RightTree = rightChild.LeftTree
	if rightChild.LeftTree != nil {
		rightChild.LeftTree.Parent = node
	}
	rightChild.Parent = node.Parent
	if node.Parent == nil {
		t.Root = rightChild
	} else if node.Parent.RightTree == node {
		node.Parent.RightTree = rightChild
	} else if node.Parent.LeftTree == node {
		node.Parent.LeftTree = rightChild
	}
	rightChild.LeftTree = node
	node.Parent = rightChild

}

func (t *AVLTree) rightRotation(node *AVLTreeNode) {
	if node == nil {
		return
	}

	leftChild := node.LeftTree
	node.LeftTree = leftChild.RightTree
	if leftChild.RightTree != nil {
		leftChild.RightTree.Parent = node
	}
	leftChild.Parent = node.Parent
	if node.Parent == nil {
		t.Root = leftChild
	} else if node.Parent.RightTree == node {
		node.Parent.RightTree = leftChild
	} else if node.Parent.LeftTree == node {
		node.Parent.LeftTree = leftChild
	}
	leftChild.RightTree = node
	node.Parent = leftChild

}

// 调整树的结构
func (t *AVLTree) fixAfterInsertion(node *AVLTreeNode, isLeft bool) {
	if isLeft {
		leftChild := node.LeftTree
		if leftChild.LeftTree != nil {
			//右旋
			t.rightRotation(node)
		} else if leftChild.RightTree != nil {
			//左右旋
			t.leftRotation(leftChild)
			t.rightRotation(node)
		}
		//return node
	}
	rightChild := node.RightTree
	if rightChild.RightTree != nil { //左旋
		t.leftRotation(node)
	} else if rightChild.LeftTree != nil { //右左旋
		t.rightRotation(rightChild)
		t.leftRotation(node)
	}
	//return node
}

func (t *AVLTree) rebuild(node *AVLTreeNode) {
	for node != nil {
		if calcBalanceFactor(node) > 1 {
			//左子树高
			t.fixAfterInsertion(node, true)
		} else if calcBalanceFactor(node) < -1 {
			t.fixAfterInsertion(node, false)
		}
		node = node.Parent
	}
}

func (t *AVLTree) putData(node *AVLTreeNode, data int) bool {
	if node == nil {
		node = buildNode(data)
		t.Root = node
		return true
	}
	sub := 0
	var p *AVLTreeNode = nil
	temp := node
	for temp != nil {
		p = temp
		sub = temp.Data - data
		if sub < 0 {
			temp = temp.RightTree
		} else if sub > 0 {
			temp = temp.LeftTree
		} else {
			return false
		}
	}
	if sub < 0 {
		p.RightTree = buildNodeByDataAndParent(data, p)
	} else if sub > 0 {
		p.LeftTree = buildNodeByDataAndParent(data, p)
	}
	t.rebuild(p) //平衡二叉树的方法
	return true
}

func (t *AVLTree) Put(data int) {
	if t.Root == nil {
		t.Root = buildNode(data)
		return
	}
	t.putData(t.Root, data)
}

func (t *AVLTree) MidOrderErgodic() {
	midOrderErgodic(t.Root)
}

func midOrderErgodic(node *AVLTreeNode) {
	if node != nil {
		midOrderErgodic(node.LeftTree)
		fmt.Println(node.Data)
		midOrderErgodic(node.RightTree)
	}
}
