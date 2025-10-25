func sortArrayByParity(nums []int) []int {
    var even []int
    var odd []int

    for _, num := range nums {
        if num%2 == 0 {
            even = append(even, num)
        } else {
            odd = append(odd, num)
        }
    }

    return append(even, odd...)
}
