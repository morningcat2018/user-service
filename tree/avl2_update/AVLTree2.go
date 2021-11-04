package avl2_update

import "fmt"

const LH = 1  /*  左高 */
const EH = 0  /*  等高 */
const RH = -1 /*  右高 */

type AvlTreeNode struct {
	data           int
	bf             int /*  结点的平衡因子 */
	lchild, rchild *AvlTreeNode
}

type BiTree *AvlTreeNode

type Status bool

func R_Rotate(p *BiTree) {
	L := (*p).lchild
	(*p).lchild = L.rchild
	L.rchild = *p
	*p = L
}

func L_Rotate(p *BiTree) {
	R := (*p).rchild
	(*p).rchild = R.lchild
	R.lchild = *p
	*p = R
}

func LeftBalance(T *BiTree) {
	L := (*T).lchild
	switch L.bf {
	case LH:
		(*T).bf = EH
		L.bf = EH
		R_Rotate(T)
		break
	case RH:
		Lr := L.rchild
		switch Lr.bf {
		case LH:
			(*T).bf = RH
			L.bf = EH
			break
		case EH:
			(*T).bf = EH
			L.bf = EH
			break
		case RH:
			(*T).bf = EH
			L.bf = LH
			break
		}
		Lr.bf = EH
		L_Rotate((*BiTree)(&(*T).lchild))
		R_Rotate(T)
	}
}

func RightBalance(T *BiTree) {
	R := (*T).rchild
	switch R.bf {
	case RH:
		(*T).bf = EH
		R.bf = EH
		L_Rotate(T)
		break
	case LH:
		Rl := R.lchild
		switch Rl.bf {
		case RH:
			(*T).bf = LH
			R.bf = EH
			break
		case EH:
			(*T).bf = EH
			R.bf = EH
			break
		case LH:
			(*T).bf = EH
			R.bf = RH
			break
		}
		Rl.bf = EH
		R_Rotate((*BiTree)(&(*T).rchild))
		L_Rotate(T)
	}
}

func InsertAVL(T *BiTree, data int, taller *Status) bool {
	if *T == nil {
		M := &AvlTreeNode{data, EH, nil, nil}
		*T = M
		*taller = true
		return true
	}
	if (*T).data == data {
		*taller = false
		return false
	}

	if data < (*T).data {
		flag := InsertAVL((*BiTree)(&(*T).lchild), data, taller)
		if !flag {
			return false
		}
		if *taller {
			switch (*T).bf {
			case LH:
				LeftBalance(T)
				*taller = false
				break
			case EH:
				(*T).bf = LH
				*taller = true
				break
			case RH:
				(*T).bf = EH
				*taller = false
				break
			}
		}
	} else {
		flag := InsertAVL((*BiTree)(&(*T).rchild), data, taller)
		if !flag {
			return false
		}

		if *taller {
			switch (*T).bf {
			case LH:
				(*T).bf = EH
				*taller = false
				break
			case EH:
				(*T).bf = RH
				*taller = true
				break
			case RH:
				RightBalance(T)
				*taller = false
				break
			}
		}
	}
	return true
}

func InOrderTraverse(t *AvlTreeNode) {
	if t != nil {
		InOrderTraverse(t.lchild)
		fmt.Printf("%d  ", t.data)
		InOrderTraverse(t.rchild)
	}
}

func RunTest() {
	a := [...]int{3, 2, 1, 7, 9, 6, 4, 5, 8, 10}
	var T BiTree = nil
	var taller Status = false
	for _, v := range a {
		InsertAVL(&T, v, &taller)
	}
	InOrderTraverse(T)
}
