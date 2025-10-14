func findDuplicates(nums []int) []int {
    seen := make(map[int]bool) 
    result := make([]int, 0)
    
    for _, num := range nums {
       
        if seen[num] {
            result = append(result, num) 
        } else {
            seen[num] = true 
        }
    }
    
    return result
}