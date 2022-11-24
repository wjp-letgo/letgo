package lib

import (
	"fmt"
	"math/rand"
	"time"
)

//随机相关的操作
//获取一个区间内的随机数
//参数min：最小值[包含最小值]
//参数max：最大值[包含最大值]
//返回最小值和最大值之间的数
func Rand(min int, max int, i int) int {
	max = max + 1
	s := rand.NewSource(time.Now().UnixNano() + int64(i))
	r := rand.New(s)
	return r.Intn(max-min) + min

}

//获取一个区间内的随机数
//参数min：最小值
//参数max：最大值
//返回最小值和最大值之间的数
func RandFloat(min float64, max float64) float64 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return (r.Float64()*(max-min) + min)
}

//获取指定长度的字符
//参数length：要获取字符的长度
//返回随机指定长度的字符串
func RandChar(length int) string {
	chr := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var s string = ""
	for i := 0; i < length; i++ {
		n := Rand(0, len(chr)-1, i)
		s += string(chr[n])
	}
	return s
}

//获取指定长度的字符
//参数length：要获取字符的长度
//返回随机指定长度的字符串
func RandCharNoNumber(length int) string {
	return RandCharNoNumberEx(length, 0)
}

//获取指定长度的字符
//参数length：要获取字符的长度
//返回随机指定长度的字符串
func RandCharNoNumberEx(length int, r int) string {
	chr := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	var s string = ""
	for i := 0; i < length; i++ {
		n := Rand(0, len(chr)-1, i+r)
		s += string(chr[n])
	}
	return s
}


//获取指定长度的字符
//参数length：要获取字符的长度
//返回随机指定长度的字符串
func RandLChar(length int) string {
	chr := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	var s string = ""
	for i := 0; i < length; i++ {
		n := Rand(0, len(chr)-1, i)
		s += string(chr[n])
	}
	return s
}

//获取指定长度的字符
//参数length：要获取字符的长度
//返回随机指定长度的字符串
func RandUChar(length int) string {
	chr := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	var s string = ""
	for i := 0; i < length; i++ {
		n := Rand(0, len(chr)-1, i)
		s += string(chr[n])
	}
	return s
}

//获取指定长度的数字
//参数length：要获取字符的长度
//返回随机指定长度的字符串
func RandNumber(length int) string {
	chr := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var s string = ""
	for i := 0; i < length; i++ {
		n := Rand(0, len(chr)-1, i)
		s += string(chr[n])
	}
	return s
}

//数组随机打乱
func RandIntArray(list []int) {
	l := len(list)
	for i := 0; i < l; i++ {
		k1 := Rand(0, l-1, i)
		k2 := Rand(0, l-1, i+1)
		tmp := list[k1]
		list[k1] = list[k2]
		list[k2] = tmp
	}
}

//UniqueCode 根据ID生成唯一码,最小长度为minLen
func UniqueCode(id int64, minLen int) string {
	g := "B"
	var r map[string][]string = map[string][]string{
		"0": []string{
			"A",
			"C",
		},
		"1": []string{
			"D",
			"E",
		},
		"2": []string{
			"F",
			"G",
			"Y",
		},
		"3": []string{
			"H",
			"I",
		},
		"4": []string{
			"Z",
			"K",
			"X",
		},
		"5": []string{
			"L",
			"M",
			"W",
		},
		"6": []string{
			"N",
			"O",
		},
		"7": []string{
			"P",
			"Q",
			"Z",
		},
		"8": []string{
			"R",
			"S",
			"V",
		},
		"9": []string{
			"T",
			"U",
		},
	}
	ids := fmt.Sprintf("%d", id)
	idsR := []rune(ids)
	idsRL := len(idsR)
	s := ""
	if idsRL >= minLen {
		for key, v := range idsR {
			k := string(v)
			m := r[k]
			max := len(m)
			i := Rand(0, max-1, key)
			s += r[k][i]
		}
	} else {
		//不用补部分
		for key, v := range idsR {
			k := string(v)
			m := r[k]
			max := len(m)
			i := Rand(0, max-1, key)
			s += r[k][i]
		}
		//位数不够补
		s += g
		//剩余补位minLen-(idsRL+1)
		for i := 0; i < minLen-(idsRL+1); i++ {
			k := fmt.Sprintf("%d", Rand(0, 9, i))
			m := r[k]
			max := len(m)
			i := Rand(0, max-1, i)
			s += r[k][i]
		}
	}
	return s
}
