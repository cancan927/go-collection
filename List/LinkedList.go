package List

import (
	"errors"
	"reflect"
)

type Node[T any] struct {
	//节点的前驱和后继
	prev, next *Node[T]

	value T
}

//LinkedList 采用带头节点的双循环链表
type LinkedList[T any] struct {
	//根节点
	root Node[T]
	//链表长度
	len int
}

//Init 初始化链表或重置链表
func (l *LinkedList[T]) init() *LinkedList[T] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

//NewLinkedList 创建一个LinkedList
func NewLinkedList[T any]() *LinkedList[T] {
	return new(LinkedList[T]).init()
}

//Size 返回链表长度
func (l *LinkedList[T]) Size() int {
	return l.len
}

//IsEmpty 判断链表是否为空
func (l *LinkedList[T]) IsEmpty() bool {
	return l.len == 0
}

//Clear 清空链表
func (l *LinkedList[T]) Clear() {
	l.init()
}

//Get 根据index获取对应的值
func (l *LinkedList[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.len {
		var t T
		return t, errors.New("index out of bounds")
	}

	p := &l.root
	if index < l.len>>1 { //index小于链表长度一半，从头节点先后遍历
		for i := 0; i <= index; i++ {
			p = p.next
		}
	} else {
		for i := l.len - 1; i >= index; i-- { //index大于链表长度一半，从尾部向前遍历
			p = p.prev
		}
	}
	return p.value, nil
}

//Contains 判断链表是否包含这个元素
func (l *LinkedList[T]) Contains(val T) bool {
	if l.IsEmpty() {
		return false
	}
	p := &l.root
	for i := 0; i < l.len; i++ {
		p = p.next
		if reflect.DeepEqual(p.value, val) {
			return true
		}
	}
	return false
}

//Set 修改目标位置的值
func (l *LinkedList[T]) Set(index int, newVal T) error {
	if index < 0 || index >= l.len {
		return errors.New("index out of bounds")
	}
	p := &l.root
	for i := 0; i <= index; i++ {
		p = p.next
	}
	p.value = newVal
	return nil
}

//Insert 在指定位置插入元素
func (l *LinkedList[T]) Insert(index int, newVal T) error {
	if index < 0 || index > l.len {
		return errors.New("index out of bounds")
	}

	if index == l.len {
		l.Append(newVal)
		return nil
	}
	p := &l.root
	for i := 0; i <= index; i++ {
		p = p.next //index ==0 时，p指向首节点
	}
	//把新元素插入到index位置元素的前面
	e := &Node[T]{value: newVal}
	//新节点前驱指向老节点的前驱
	e.prev = p.prev //index == 0 时，p.prev 指向root   <--E
	//新节点后继指向老节点
	e.next = p //  E-->
	//老前驱节点的后继指向新节点
	e.prev.next = e //root.next = e    -->E
	//老节点前驱指向新节点
	e.next.prev = e //       E<--
	l.len++
	return nil
}

//Append 在末尾添加新元素
func (l *LinkedList[T]) Append(newVal T) {
	p := &Node[T]{value: newVal}
	//新节点后继指向根节点
	p.next = &l.root
	//新节点的前驱指向之前的尾节点
	p.prev = l.root.prev
	//旧尾节点的后继指向新节点
	l.root.prev.next = p
	//根节点的前驱指向新节点
	l.root.prev = p
	l.len++
}

//Delete 删除指定位置的值
func (l *LinkedList[T]) Delete(index int) error {
	if index < 0 || index >= l.len {
		return errors.New("index out of bounds")
	}
	p := &l.root
	for i := 0; i <= index; i++ {
		p = p.next
	}
	p.prev.next = p.next
	p.next.prev = p.prev
	l.len--
	return nil
}
