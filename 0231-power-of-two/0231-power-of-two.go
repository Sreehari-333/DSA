// func isPowerOfTwo(n int) bool {
    
//     for i:=0 ; i <= n ; i++{
//         if int(math.Pow(2,float64(i))) == n {
//             return true
//         }
//     }
//     return false
// }
func isPowerOfTwo(n int) bool {
    return n > 0 && (n & (n - 1)) == 0
}