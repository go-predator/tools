/*
 * @Author:    thepoy
 * @Email:     thepoy@163.com
 * @File Name: string.go
 * @Created:   2021-11-24 19:58:22
 * @Modified:  2022-05-24 10:15:27
 */

package tools

import "strings"

// Trim 删除目标字符串两侧的无用字符
func Trim(s_ string, chars_ string) string {
	s, chars := []rune(s_), []rune(chars_)
	length := len(s)
	max := len(s) - 1
	l, r := true, true //标记当左端或者右端找到正常字符后就停止继续寻找
	start, end := 0, max
	tmpEnd := 0
	charset := make(map[rune]bool) //创建字符集&#xff0c;也就是唯一的字符&#xff0c;方便后面判断是否存在
	for i := 0; i < len(chars); i++ {
		charset[chars[i]] = true
	}
	for i := 0; i < length; i++ {
		if _, exist := charset[s[i]]; l && !exist {
			start = i
			l = false
		}
		tmpEnd = max - i
		if _, exist := charset[s[tmpEnd]]; r && !exist {
			end = tmpEnd
			r = false
		}
		if !l && !r {
			break
		}
	}
	if l && r { // 如果左端和右端都没找到正常字符&#xff0c;那么表示该字符串没有正常字符
		return ""
	}
	return string(s[start : end+1])
}

// Trim 删除两则空白字符，包插空格、换行、制表、全角空格
func Strip(src string) string {
	return Trim(src, " \n\t 　")
}

// ReplaceInvalidSpaces 替换html页面中各种奇葩的空白符为空格
func ReplaceInvalidSpaces(src string) string {
	src = strings.ReplaceAll(src, " ", " ")
	src = strings.ReplaceAll(src, "　", " ")

	return src
}

// RemoveInvalidChars 删除单句中的无效字符，如一些无意义的特殊字符
func RemoveInvalidChars(src string, chars ...string) string {
	if len(chars) == 0 {
		return src
	}

	for _, char := range chars {
		src = strings.ReplaceAll(src, char, "")
	}

	return ReplaceInvalidSpaces(src)
}
