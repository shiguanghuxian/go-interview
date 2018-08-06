package main

/* 单向链表 */

import (
	"fmt"
	"log"
)

// 定义链表节点
type node struct {
	Data int
	Next *node
}

func main() {
	// 系统日志显示文件和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	head := &node{0, nil}
	link := newLink(head, 5)
	// 打印链表
	// js, _ := json.Marshal(link)
	// log.Println(string(js))
	printLink(link)

	// 反转链表
	link2 := reverseLink2(link)
	// 打印
	printLink(link2)
}

// 创建链表
func newLink(head *node, length int) *node {
	if length < 1 {
		return head
	}
	// 循环创建从链表最后一个开始创建链表
	for i := length - 1; i > 0; i-- {
		p := &node{i, nil}
		// 将head的下级放到新创建的节点下级，再将新建节点放到head下级
		// 每创建一个元素，将元素插入head和之后节点之间
		p.Next = head.Next
		head.Next = p
	}
	return head
}

// 打印链表
func printLink(head *node) {
	linkStr := ""
	for p := head; p != nil; p = p.Next {
		if linkStr == "" {
			linkStr = fmt.Sprintf(" %d", p.Data)
		} else {
			linkStr = fmt.Sprintf("%s -> %d", linkStr, p.Data)
		}
	}
	log.Println(linkStr)
}

// 反转链表 - 单向链表
func reverseLink(head *node) *node {
	if head == nil || head.Next == nil {
		return head
	}
	var reversedHead *node
	var p *node
	// reversedHead 会一直将自己和自己下一级节点交换
	// p 指向下一个节点
	reversedHead = head
	head = head.Next
	reversedHead.Next = nil
	p = head.Next
	for head != nil {
		head.Next = reversedHead
		reversedHead = head
		head = p
		if p != nil {
			p = p.Next
		}
	}
	return reversedHead
}

// 反转链表2 - 单向链表
func reverseLink2(head *node) *node {
	if head == nil || head.Next == nil {
		return head
	}

	// 将所有节点放入数组
	links := make([]*node, 0)
	for p := head; p != nil; p = p.Next {
		links = append(links, p)
	}
	links = append(links, nil)

	// 和reverseLink原理相同
	var reversedHead *node = links[0]
	head = links[1]
	reversedHead.Next = nil

	for i := 2; i < len(links); i++ {
		head.Next = reversedHead
		reversedHead = head
		if links[i] != nil {
			head = links[i]
		}
	}

	return reversedHead
}
