package tree

import (
	"fmt"
	"math"
)

type Color int

const (
	BLACK Color = 0
	RED   Color = 1
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	color Color
	p     *TreeNode //parent
}

type Tree struct {
	Root *TreeNode
}

func GenerateRadBlackTree(nums []int) Tree {
	t := Tree{}
	for _, v := range nums {
		t.InsertRedBlackTree(v)
	}
	return t
}

func (t *Tree) InsertRedBlackTree(n int) {
	//four case
	tn := &TreeNode{
		Val:   n,
		color: RED,
	}
	if t.Root == nil {
		tn.color = BLACK
		t.Root = tn
		return
	}
	t.Insert(tn)
	t.ReBalance(tn)

}

func (t *Tree) Insert(tn *TreeNode) {
	nowNode := t.Root
	findInsertNode(nowNode, tn)
}

func findInsertNode(nowNode, tn *TreeNode) {
	if nowNode.Val > tn.Val {
		if nowNode.Left == nil {
			nowNode.Left = tn
			tn.p = nowNode
		} else {
			findInsertNode(nowNode.Left, tn)
		}
	} else {
		if nowNode.Right == nil {
			nowNode.Right = tn
			tn.p = nowNode
		} else {
			findInsertNode(nowNode.Right, tn)
		}
	}
}

func (t *Tree) ReBalance(tn *TreeNode) {
	p := tn.p
	if p == nil {
		tn.color = BLACK
		t.Root = tn
		return
	}
	u := findUncle(tn)
	g := p.p
	if p.color == RED && (u == nil || u.color == BLACK) {
		//line
		if g.Left == p && p.Left == tn {
			t.rightRotate(g)
			p.color = BLACK
			g.color = RED
		} else if g.Right == p && p.Right == tn {
			t.leftRotate(g)
			p.color = BLACK
			g.color = RED
		} else if g.Left == p && p.Right == tn {
			t.leftRotate(g)
			t.ReBalance(g.p)
		} else if g.Right == p && p.Left == tn {
			t.rightRotate(g)
			t.ReBalance(g.p)
		}
	} else if p.color == RED && u.color == RED {
		p.color = BLACK
		if u != nil {
			u.color = BLACK
		}
		g.color = RED
		t.ReBalance(g)
	}
}

func changeColor(tn *TreeNode) {
	tn.color = tn.color ^ BLACK
}

func findUncle(tn *TreeNode) *TreeNode {
	p := tn.p
	if p == nil {
		return nil
	} else {
		g := p.p
		if g == nil {
			return nil
		}
		if g.Left == p {
			return g.Right
		} else {
			return g.Left
		}
	}
}

func (t *Tree) leftRotate(tn *TreeNode) {
	oldRight := tn.Right
	tn.Right = oldRight.Left
	if oldRight.Left != nil {
		oldRight.Left.p = tn
	}
	oldRight.p = tn.p
	if tn.p == nil {
		oldRight.p = nil
		t.Root = oldRight
	} else if tn == tn.p.Left {
		tn.p.Left = oldRight
	} else {
		tn.p.Right = oldRight
	}
	oldRight.Left = tn
	tn.p = oldRight
}

func (t *Tree) rightRotate(tn *TreeNode) {
	oldLeft := tn.Left
	tn.Left = oldLeft.Right
	if oldLeft.Right != nil {
		oldLeft.Left.p = tn
	}
	oldLeft.p = tn.p
	if tn.p == nil {
		oldLeft.p = nil
		t.Root = oldLeft
	} else if tn == tn.p.Left {
		tn.p.Left = oldLeft
	} else {
		tn.p.Right = oldLeft
	}
	oldLeft.Right = tn
	tn.p = oldLeft
}

func (t Tree) Print(ch chan *TreeNode) {
	dfs(t.Root, ch)
	close(ch)
}

func dfs(t *TreeNode, ch chan *TreeNode) {
	if t == nil {
		return
	}

	dfs(t.Left, ch)
	ch <- t
	dfs(t.Right, ch)
}

func getTreeLevel(tn *TreeNode) int {
	if tn == nil {
		return 0
	}
	l := getTreeLevel(tn.Left)
	r := getTreeLevel(tn.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

func (t Tree) PrettyPrint() {
	level := getTreeLevel(t.Root)
	res := make([][]byte, level)
	tmp := make([][]*TreeNode, level)
	colNum := math.Pow(2, float64(level - 1))
	for i := range res {
		res[i] = make([]byte, int(colNum * 2))
		//tmp[i] = make([]*TreeNode, 1)
	}

	for _, v := range tmp {
		for _, vv := range v {
			fmt.Print(vv)
		}
		fmt.Println("")
	}
}


func (t Tree) LevelPrint() {
	ch := make(chan *TreeNode, 16)
	ch <- t.Root

	for {
		var tmp []*TreeNode
		for len(ch) != 0 {
			tn := <- ch
			tmp = append(tmp, tn)
			fmt.Print(tn.Val, "      ")
		}
		fmt.Println()
		if tmp == nil {
			break
		}
		//fmt.Println(len(tmp))

		for _, v := range tmp {
			if v.Left != nil {
				ch <- v.Left
			}
			if v.Right != nil {
				ch <- v.Right
			}
		}
	}
}

