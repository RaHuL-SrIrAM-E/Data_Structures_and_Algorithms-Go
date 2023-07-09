package main
import "fmt"

type node struct{
	data int
	next *node
}
type LinkedList struct{
	head *node
}
func print(list *LinkedList){
	current:=list.head
	for current !=nil {
		fmt.Printf("%v",current.data)
		fmt.Printf("->")
		current=current.next
	} 
	fmt.Println()
}
func (list *LinkedList) InsertFirst(value int){
	newNode:=&node{data: value}
	if list.head==nil {
		list.head=newNode
		return
	} else{
		newNode.next=list.head
		list.head=newNode
	}
}
func (list *LinkedList) InsertLast(value int){
	newNode:=&node{data:value}
	current:=list.head
	for current.next!=nil{
		current=current.next
	}
	current.next=newNode
}
func GetSize(list *LinkedList)int{
	current:=list.head
	var count int=0
	for current!=nil{
		count++
		current=current.next
	}
	return count
}
func (list *LinkedList) InsertAtIndex(index int, value int){
	var count int = 0
	current:=list.head
	for count!=index{
		count++
		current=current.next
	}
	newNode:=&node{data:value}
	newNode.next=current.next
	current=newNode
}
func (list *LinkedList) RemoveLast(){
	if list.head==nil{
		return
	}
	current:=list.head
	for current.next!=nil{
		if current.next.next!=nil {
			current=current.next
		}else{
			current.next=nil
		}
	}
}
func (list *LinkedList) GetFirst()(int,bool){
	if list.head==nil {return 0,false}
	return list.head.data,true
}
func (list *LinkedList) GetAtIndex(value int)(int,bool){
	if list.head==nil {return 0,false}
	if value<0 {return -1,false}
	current:=list.head
	var count int=0
	for current.next!=nil {
		if current.next.next!=nil {
			if count!=value {
				count++
				current=current.next
			}else{
				return current.data,true
			}
		}
	}
	return 0,false
}
func (list *LinkedList) RemoveAtIndex(i int) {
	if list.head == nil {return}
	if i < 0 {
		return
	}
	if i == 0 {
		list.head = list.head.next
		return
	}
	current := list.head
	for u := 1; u < i; u++ {
		if current.next.next == nil {
			return
		}
		current = current.next
	}
	current.next = current.next.next
}
func (list *LinkedList) Search(value int)bool{
	if list.head==nil {return false}
	current:=list.head
	for current!=nil {
		if current.data==value {
			return true
		}else{
			current=current.next
		}
	}
	return false
}
func (list *LinkedList) ListToArray() []int{
	var array []int
	current:=list.head
	for current!=nil{
		array=append(array, current.data)
		current=current.next
	}
	return array
}
func main(){
	//program for linked list
	list := &LinkedList{}
	print(list)
	list.InsertFirst(2)
	print(list)
	list.InsertFirst(3)
	print(list)
	list.InsertLast(5)
	print(list)
	count:=GetSize(list)
	fmt.Println(count)
	list.RemoveLast()
	print(list)
	fmt.Println(list.GetFirst())
	list.InsertFirst(4)
	list.InsertFirst(16)
	list.InsertFirst(20)
	print(list)
	fmt.Println(list.GetAtIndex(2))
	list.RemoveAtIndex(2)
	print(list)
	fmt.Println(list.Search(4))
}
