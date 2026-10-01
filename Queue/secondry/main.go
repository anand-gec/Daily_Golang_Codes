// Queue implemented in LinkedList
package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func Enqueue(val int, head *Node) *Node {
	var temp = &Node{val:val,next:nil}
	if head == nil {
		head = temp
		return head
	}
	t := head
	for t.next != nil {
		t = t.next
	}
	t.next = temp
	return head
}
func Display(head *Node) {
	for head != nil {
		fmt.Print(" -> ", head.val)
		head = head.next
	}
	fmt.Println("")
}

func Front(head *Node)int{
if head==nil{
	return 0
}
return head.val
}

func Rare(head *Node)int{
if head==nil{
	fmt.Println("not Contain any Value")
	return 0
}
var num=0
t:=head
for t!=nil{
	if t.next==nil{
		num=t.val
	}
	t=t.next
}
return num
}
func Dequeue(head *Node) *Node {
	if head == nil || head.next == nil {
		fmt.Println("Queue is Vacant... ")
		return nil
	}
	fmt.Println("Delete Value is ",head.val)
	nextNode := head.next
	head = nil
	return nextNode
}
func Length(head *Node)int{
	if head==nil{
		return 0
	}
	var len=0
	for head!=nil{
		len++
		head=head.next
	}
	return len
}

func main() {
	var head *Node
	head = Enqueue(2, head)
	head = Enqueue(4, head)
	head = Enqueue(6, head)
	head = Enqueue(8, head)
	head = Dequeue(head)
	head = Dequeue(head)
	// head = Dequeue(head)
	Display(head)
	fmt.Println("Front Value is :",Front(head))
	fmt.Println("Rare Value is :",Rare(head))
	fmt.Println("Length of Queue is",Length(head))
}
