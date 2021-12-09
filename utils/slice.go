package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func SliceAtoi(slice []string) []int {
	l := []int{}
	for _, s := range slice {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		l = append(l, i)
	}
	return l
}

func IntJoin(slice []int, sep string) string {
	ss := []string{}
	for _, s := range slice {
		ss = append(ss, strconv.Itoa(s))
	}
	return strings.Join(ss, ",")
}

func FloatJoin(slice []float64, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	ss := []string{}
	for _, s := range slice {
		ss = append(ss, fmt.Sprintf("%v", s))
	}
	return strings.Join(ss, sep)
}

func IntSplit(str string, sep string) []int {
	ss := strings.Split(str, sep)
	vs := []int{}
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		if i == 0 {
			continue
		}
		vs = append(vs, i)
	}
	return vs
}

func InArray(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func IntersectArrays(arrs [][]int) []int {
	a := []int{}
	for i, arr := range arrs {
		if i == 0 {
			a = arr
		}
		a = IntersectArray(a, arr)
	}
	return a
}

// 交集
func IntersectArray(arr1 []int, arr2 []int) []int {
	newArr := []int{}
	m := map[int]int{}
	for _, v := range arr1 {
		m[v] = 1
	}

	for _, v := range arr2 {
		if _, ok := m[v]; ok {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

// 差集, 在arr1中，不在arr2中
func DiffArray(arr1 []int, arr2 []int) []int {
	newArr := []int{}
	m := map[int]int{}
	for _, v := range arr2 {
		m[v] = 1
	}

	for _, v := range arr1 {
		if _, ok := m[v]; !ok {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

// 差集, 在arr1中，不在arr2中
func DiffArrayString(arr1 []string, arr2 []string) []string {
	newArr := []string{}
	m := map[string]int{}
	for _, v := range arr2 {
		m[v] = 1
	}

	for _, v := range arr1 {
		if _, ok := m[v]; !ok {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

func KeysMap(mm map[int]int) []int {
	ids := []int{}
	for m, _ := range mm {
		ids = append(ids, m)
	}
	return ids
}

func UniqArray(arr []int) []int {
	aa := []int{}
	m := map[int]int{}
	for _, a := range arr {
		if _, ok := m[a]; ok {
			continue
		}
		aa = append(aa, a)
		m[a] = 1
	}
	return aa
}

func UniqArrayString(arr []string) []string {
	aa := []string{}
	m := map[string]int{}
	for _, a := range arr {
		if _, ok := m[a]; ok {
			continue
		}
		aa = append(aa, a)
		m[a] = 1
	}
	return aa
}

func TranslateArray(d interface{}) []string {
	ret := []string{}
	for _, a := range d.([]interface{}) {
		ret = append(ret, a.(string))
	}

	return ret
}

func InStrArray(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func IsSameArray(a1, a2 []string) bool {
	if len(a1) != len(a2) {
		return false
	}
	if len(a1) == 0 {
		return true
	}
	check := map[string]int{}

	for _, a := range a1 {
		if _, ok := check[a]; ok {
			check[a]++
		} else {
			check[a] = 0
		}
	}
	for _, a := range a2 {
		if _, ok := check[a]; ok {
			check[a]--
		} else {
			return false
		}
	}
	for _, a := range check {
		if a != 0 {
			return false
		}
	}
	return true
}
