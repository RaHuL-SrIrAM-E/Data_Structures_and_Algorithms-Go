//program to reverse a string by removing repeated characters
package main
import "fmt"

func funmap(s string)string{
	var result string
	m:=make(map[string]int)
	for i:=len(s)-1; i>=0;i--{
		if _,ch:=m[string(s[i])] ; !ch{
			result=result+string(s[i])
			m[string(s[i])]=1
		} else {
			m[string(s[i])]++
		}
	}
	return result
}
func main(){
	var s string
	fmt.Scanln(&s)
	fmt.Println(s)
	result:=funmap(s)
	fmt.Println(result)
}
