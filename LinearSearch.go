package main
import "fmt"

func LinearSearch(array [5]int, value int)int{
	for i:=0;i<len(array);i++ {
		if array[i]==value {return i}
	}
	return -1
}
func main(){
	b := [5]int{1, 2, 3, 4, 5}
	c:=4
	fmt.Println(LinearSearch(b,c))
}
