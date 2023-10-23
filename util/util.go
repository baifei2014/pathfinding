// @Author jianglonghao
// @Date 2023/10/23
// @Description

package util

func Reverse[T any](s []T) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
}
