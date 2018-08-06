package main

/* 双向链表 */

import (
	"fmt"
	"log"
)

type node struct {
	NextNode *node
	LastNode *node
	Data     int
}

func main() {
	// 系统日志显示文件和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	head := new(node)
	head.Data = 1
	link := newLink(head, 10)
	// 反转前
	printLink(link)
	// 反转后
	link = reverseLink2(link)
	printLink(link)

	// 再次反转
	link = reverseLink2(link)
	printLink(link)
}

// 创建链表
func newLink(head *node, length int) *node {
	if length < 1 {
		return head
	}
	head.LastNode = nil

	for i := length - 1; i > 1; i-- {
		p := &node{
			NextNode: head.NextNode,
			LastNode: head,
			Data:     i,
		}
		head.NextNode = p
	}
	// 反序定位
	var last *node
	for p := head; p != nil; p = p.NextNode {
		p.LastNode = last
		last = p
	}

	return head
}

// 打印链表
func printLink(head *node) {
	linkStr := ""
	for p := head; p != nil; p = p.NextNode {
		if linkStr == "" {
			linkStr = fmt.Sprintf(" %d", p.Data)
		} else {
			linkStr = fmt.Sprintf("%s -> %d", linkStr, p.Data)
		}
	}

	log.Println(linkStr)
}

// 反转
func reverseLink(head *node) *node {
	if head == nil || head.NextNode == nil {
		return head
	}
	nodes := make([]*node, 0)
	for p := head; p != nil; p = p.NextNode {
		nodes = append(nodes, p)
	}
	for _, p := range nodes {
		// log.Println(p.Data)
		p.LastNode, p.NextNode = p.NextNode, p.LastNode
	}
	return nodes[len(nodes)-1]
}

// 反转2
func reverseLink2(head *node) *node {
	if head == nil || head.NextNode == nil {
		return head
	}

	var reversedHead *node
	var p *node

	reversedHead = head
	head = head.NextNode
	reversedHead.NextNode = nil
	p = head.NextNode
	for head != nil {
		head.NextNode = reversedHead
		reversedHead = head
		head = p
		if p != nil {
			p = p.NextNode
			if p != nil {
				p.LastNode = head
			}
		}
	}

	return reversedHead
}
