package main


import "fmt"


// "how do you copy a slice, map, and an interface?"

func main(){
// copy a slice
slice1 := []int{1,2,4}
slice2 := make([]int, len(slice1))
fmt.Println(slice1, slice2)
copy(slice2, slice1)
fmt.Println(slice1, slice2)

// copy a map -- there might be a better way with builtins
map1, map2 := make(map[int]int), make(map[int]int)
map1 = map[int]int{1:2,3:4,5:6}
fmt.Println(map1,map2)
for key,value := range map1{
    map2[key] = value
}
fmt.Println(map1,map2)







}
