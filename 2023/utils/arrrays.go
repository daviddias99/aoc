package	utils

func AllZeros(s []int) bool {
    for _, v := range s {
        if v != 0 {
            return false
        }
    }
    return true
}