package avl2

import "fmt"

const LH = 1  /*  左高 */
const EH = 0  /*  等高 */
const RH = -1 /*  右高 */

type AvlTreeNode struct {
	data           int
	bf             int /*  结点的平衡因子 */
	lchild, rchild *AvlTreeNode
}

type AvlTree struct {
	root *AvlTreeNode
}

type Status bool

//参考实现：https://www.cnblogs.com/polly333/p/4798944.html
//对以p为根的二叉排序树作右旋处理
//处理之后p指向新的树根结点，即旋转处理之前的左子树的根结点
//右旋-顺时针旋转(如LL型就得对根结点做该旋转)
func (t *AvlTree) R_Rotate(p *AvlTreeNode) *AvlTreeNode {
	L := p.lchild       // L指向P的左子树根结点
	p.lchild = L.rchild // L的右子树挂接为P的左子树
	L.rchild = p
	p = L /*  P指向新的根结点 */
	return p
}

/* 对以P为根的二叉排序树作左旋处理， */
/* 处理之后P指向新的树根结点，即旋转处理之前的右子树的根结点0  */
//左旋-逆时针旋转(如RR型就得对根结点做该旋转)
func (t *AvlTree) L_Rotate(p *AvlTreeNode) *AvlTreeNode {
	R := p.rchild       // R指向P的右子树根结点
	p.rchild = R.lchild // R的左子树挂接为P的右子树
	R.lchild = p
	p = R // P指向新的根结点
	return p
}

/*  对以指针T所指结点为根的二叉树作左平衡旋转处理 */
/*  本算法结束时，指针T指向新的根结点 */
func (t *AvlTree) LeftBalance(T *AvlTreeNode) *AvlTreeNode {
	L := T.lchild /*  L指向T的左子树根结点 */
	switch L.bf {
	/* 检查T的左子树的平衡度，并作相应平衡处理 */
	case LH: /* 新结点插入在T的左孩子的左子树上，要作单右旋处理 */
		T.bf = EH
		L.bf = EH
		T = t.R_Rotate(T)
		break
	case RH: /* 新结点插入在T的左孩子的右子树上，要作双旋处理 */ //
		Lr := L.rchild /* Lr指向T的左孩子的右子树根 */
		switch Lr.bf {
		/* 修改T及其左孩子的平衡因子 */
		case LH:
			T.bf = RH
			L.bf = EH
			break
		case EH:
			T.bf = EH
			L.bf = EH
			break
		case RH:
			T.bf = EH
			L.bf = LH
			break
		}
		Lr.bf = EH
		T.lchild = t.L_Rotate(T.lchild) /* 对T的左子树作左旋平衡处理 */
		T = t.R_Rotate(T)               /* 对T作右旋平衡处理 */
	}
	return T
}

/*  对以指针T所指结点为根的二叉树作右平衡旋转处理， */
/*  本算法结束时，指针T指向新的根结点 */
func (t *AvlTree) RightBalance(T *AvlTreeNode) *AvlTreeNode {
	R := T.rchild /*  R指向T的右子树根结点 */
	switch R.bf {
	/*  检查T的右子树的平衡度，并作相应平衡处理 */
	case RH: /*  新结点插入在T的右孩子的右子树上，要作单左旋处理 */
		T.bf = EH
		R.bf = EH
		T = t.L_Rotate(T)
		break
	case LH: /*  新结点插入在T的右孩子的左子树上，要作双旋处理 */ //最小不平衡树的根结点为负，其右孩子为正
		Rl := R.lchild /*  Rl指向T的右孩子的左子树根 */
		switch Rl.bf {
		/*  修改T及其右孩子的平衡因子 */
		case RH:
			(*T).bf = LH
			R.bf = EH
			break
		case EH:
			T.bf = EH
			R.bf = EH
			break
		case LH:
			(*T).bf = EH
			R.bf = RH
			break
		}
		Rl.bf = EH
		T.rchild = t.R_Rotate(T.rchild) /*  对T的右子树作右旋平衡处理 */
		T = t.L_Rotate(T)               /*  对T作左旋平衡处理 */
	}
	return T
}

/*  若在平衡的二叉排序树T中不存在和e有相同关键字的结点，则插入一个 */
/*  数据元素为e的新结点，并返回1，否则返回0。若因插入而使二叉排序树 */
/*  失去平衡，则作平衡旋转处理，布尔变量taller反映T长高与否。 */
func (t *AvlTree) InsertAVL(T *AvlTreeNode, data int, taller *Status) (bool, *AvlTreeNode) {
	if T == nil {
		T = &AvlTreeNode{data, EH, nil, nil}
		*taller = true
		return true, T
	}
	if T.data == data {
		/*  树中已存在和e有相同关键字的结点则不再插入 */
		*taller = false
		return false, nil
	}

	if data < T.data {
		flag, P := t.InsertAVL(T.lchild, data, taller)
		T.lchild = P
		if !flag {
			return false, nil
		}
		if *taller {
			/*   已插入到T的左子树中且左子树“长高” */
			switch (*T).bf { /*  检查T的平衡度 */
			case LH: /*  原本左子树比右子树高，需要作左平衡处理 */
				T = t.LeftBalance(T)
				*taller = false
				break
			case EH: /*  原本左、右子树等高，现因左子树增高而使树增高 */
				(*T).bf = LH
				*taller = true
				break
			case RH: /*  原本右子树比左子树高，现左、右子树等高 */
				(*T).bf = EH
				*taller = false
				break
			}
		}
	} else {
		flag, P := t.InsertAVL(T.rchild, data, taller)
		T.rchild = P
		if !flag {
			return false, nil
		}

		if *taller {
			/*  已插入到T的右子树且右子树“长高” */
			switch T.bf { /*  检查T的平衡度 */
			case LH: /*  原本左子树比右子树高，现左、右子树等高 */
				T.bf = EH
				*taller = false
				break
			case EH: /*  原本左、右子树等高，现因右子树增高而使树增高  */
				T.bf = RH
				*taller = true
				break
			case RH: /*  原本右子树比左子树高，需要作右平衡处理 */
				T = t.RightBalance(T)
				*taller = !true
				break
			}
		}
	}
	return true, T
}

func InOrderTraverse(t *AvlTreeNode) {
	if t != nil {
		InOrderTraverse(t.lchild)
		fmt.Printf("%d  ", t.data)
		InOrderTraverse(t.rchild)
	}
}

func RunTest() {
	a := [...]int{17, 10, 30, 22, 40, 33, 2, 8, 6}
	t := &AvlTree{nil}
	var taller Status = false
	for _, v := range a {
		_, P := t.InsertAVL(t.root, v, &taller)
		t.root = P
	}
	InOrderTraverse(t.root)
}
