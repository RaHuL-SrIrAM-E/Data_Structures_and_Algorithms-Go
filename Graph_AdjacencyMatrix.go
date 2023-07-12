package main
import "fmt"

const n=5

type Graph struct{
	Matrix[n][n]bool
}
func (graph *Graph)print(){
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
		   fmt.Printf("%t ", graph.Matrix[i][j])
		}
		fmt.Println()
	}
}
func main(){

	graph:=Graph{}
	graph.Matrix[0][1] = true
	graph.Matrix[0][2] = true
	graph.Matrix[1][2] = true

	graph.print()
}
