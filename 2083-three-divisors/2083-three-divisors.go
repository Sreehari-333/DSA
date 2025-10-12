func isThree(n int) bool {
    numOfDivisors := 0
    for i:=1;i<=n;i++{
        if n%i==0{
            numOfDivisors++
        }
    }

    if numOfDivisors ==3 {
        return true
    }
    return false
}