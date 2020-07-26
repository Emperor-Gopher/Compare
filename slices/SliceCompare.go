package slices

import (
	"reflect"
)

func CompareSlice(x interface{}, y interface{}, order bool) string {

	if reflect.TypeOf(x).Kind() != reflect.Slice {
		panic("first param is not a slice")
	}

	if reflect.TypeOf(y).Kind() != reflect.Slice {
		panic("second param is not a slice")
	}

	s1 := reflect.ValueOf(x)
	s2 := reflect.ValueOf(y)

	l1 := s1.Len()
	l2 := s2.Len()

	slc1 := make([]interface{}, l1, l1)
	slc2 := make([]interface{}, l2, l2)

	for i := 0; i < l1; i++ {
		slc1 = append(slc1, s1.Index(i).Interface())
	}

	for i := 0; i < l1; i++ {
		slc2 = append(slc2, s2.Index(i).Interface())
	}

	var res string
	if order {
		res = OrderedCompare(slc1, slc2, cap(slc1), cap(slc1))
	} else {
		res = UnorderedCompare(slc1, slc2, cap(slc1), cap(slc1))
	}

	return res
}

func OrderedCompare(slc1 []interface{}, slc2 []interface{}, l1 int, l2 int) string {

	result := "unequal"
	if l1 == l2 {

		flag := true
		for i := 0; i < l1; i++ {
			if slc1[i] != slc2[i] {
				flag = false
				break
			}
		}
		if flag {
			result = "equal"
		}
	}

	return result
}

func UnorderedCompare(slc1 []interface{}, slc2 []interface{}, l1 int, l2 int) string {

	result := "unequal"
	if l1 == l2 {

		flag := true
		m1 := make(map[interface{}]int)
		m2 := make(map[interface{}]int)

		for i := 0; i < l1; i++ {

			if 0 == m1[slc1[i]] {
				m1[slc1[i]] = 1
			} else {
				m1[slc1[i]] = m1[slc1[i]] + 1
			}
		}

		for i := 0; i < l2; i++ {

			if 0 == m2[slc2[i]] {
				m2[slc2[i]] = 1
			} else {
				m2[slc2[i]] = m2[slc2[i]] + 1
			}
		}
		for k, v := range m1 {
			if m2[k] != v {
				flag = false
				break
			}
		}

		if flag {
			result = "equal"
		}
	}
	return result
}
