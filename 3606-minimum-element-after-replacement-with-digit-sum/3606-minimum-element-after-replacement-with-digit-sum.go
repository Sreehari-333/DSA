func minElement(nums []int) int {
    sumOfDigits := []int{}
    for _, num := range nums {
        sum := 0
        strNum := strconv.Itoa(num)
        for _, ch := range strNum {
            digit, _ := strconv.Atoi(string(ch))
            sum += digit
        }
        sumOfDigits = append(sumOfDigits, sum)
    }

    sort.Ints(sumOfDigits)
    return sumOfDigits[0]
}