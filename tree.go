package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

/* 查找二叉树最远两点距离 */

// 二叉树结构体
type tree struct {
	Value             int
	Left, Right       *tree
	MaxLeft, MaxRight int
}

var nMaxLen int = 0 // 保存最长两点距离

func main() {
	// 系统日志显示文件和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	/* 生成一个随机二叉树 */
	// 二叉树最大深度
	depth := getRandInt(10)
	// depth = 3
	// 生成
	root := new(tree)
	makeTree(root, depth)

	// 计算最远两点
	findMaxLen(root)

	// 打印二叉树为json
	js, _ := json.Marshal(root)
	log.Println(string(js))

	// 打印结果
	log.Println(nMaxLen)
}

// 生成二叉树的左右节点
func makeTree(root *tree, depth int) {
	if depth == 0 {
		return
	}
	// 左侧节点是否有值
	l := getRandInt(1000)
	if l > 500 {
		root.Left = new(tree)
		makeTree(root.Left, depth-1)
	}
	r := getRandInt(1000)
	if r > 500 {
		root.Right = new(tree)
		makeTree(root.Right, depth-1)
	}
}

// 递归查找最远距离
func findMaxLen(one *tree) {
	if one == nil {
		return
	}
	// 左侧若为nil，左侧长度为0
	if one.Left == nil {
		one.MaxLeft = 0
	} else {
		findMaxLen(one.Left)
		if one.Left.MaxLeft > one.Left.MaxRight {
			one.MaxLeft = one.Left.MaxLeft
		} else {
			one.MaxLeft = one.Left.MaxRight + 1
		}
	}
	// 右侧若为nil，左侧长度为0
	if one.Right == nil {
		one.MaxRight = 0
	} else {
		findMaxLen(one.Right)
		if one.Right.MaxLeft > one.Right.MaxRight {
			one.MaxRight = one.Right.MaxLeft
		} else {
			one.MaxRight = one.Right.MaxRight + 1
		}
	}
	// 更新最长距离
	if one.MaxLeft+one.MaxRight > nMaxLen {
		nMaxLen = one.MaxLeft + one.MaxRight
	}
}

// 随机数生成(时间为种子)
var s1 = rand.NewSource(time.Now().Unix())
var r1 = rand.New(s1)

// 获取一个随机数
func getRandInt(max int) int {
	i := r1.Intn(max)
	log.Println(i)
	return i
}
