package main
import "fmt"

func twoSum(nums []int, target int) []int {
    index := make(map[int]int)
    for i := range nums {
        j, ok := index[target-nums[i]]
        if ok {
            result := []int{j, i}
            return result
        }
        index[nums[i]] = i
    }
    result := []int{-1, -1}
    return result    
}


func main() {
    num := []int{2, 7, 11, 17}
    target := 9
    fmt.Println(twoSum(num, target))
}
