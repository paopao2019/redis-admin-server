package uitls


// 遍历 切片 求和
func Sum(arr []int) (result int) {
	for _, v := range arr {
		result += v
	}
	return result
}
