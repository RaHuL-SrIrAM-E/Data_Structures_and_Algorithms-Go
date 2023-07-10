func BinarySearch(array []int, number int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex {
		midIndex := int((maxIndex + minIndex) / 2)
		midItem := array[midIndex]
		if number == midItem {
			return midIndex
		}
		if midItem < number {
			minIndex = midIndex + 1
		} else if midItem > number {
			maxIndex = midIndex - 1
		}
	}
	return -1
}
func main(){
  array:=[5]{1,2,3,4,5}
  value:=4
  fmt.Println(BinarySearch(array,value))
}
