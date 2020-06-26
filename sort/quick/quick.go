package quick

import "fmt"

var arr = []int{50, 10, 90, 30, 70, 40, 80, 60, 20}

/*
1.选取基准值tag：选取第一个元素为基准值（每次递归调用的基准值不等）
2.从0向后查找到比tag大的数，否则i++
3.从n向前查找到比tag小的数，否则j--
4.交换两个数
5.全部交换后此时i==j,递归调用，继续比较i位置两侧位置
*/

// 第一种方式，递归传入分割后的数组
func quicksort1(s []int) {
	fmt.Println("原数组: ", arr)
	if len(s) <= 0 {
		return
	}
	tag := s[0]
	i := 0
	j := len(s) - 1
	for i < j {
		for s[j] > tag && i < j {
			j--
		}
		for s[i] < tag && i < j {
			i++
		}
		s[i], s[j] = s[j], s[i] //交换
		//fmt.Println("交换完之后的数组: ", s)
	}
	quicksort1(s[:i])   // 左子序列 排序
	quicksort1(s[i+1:]) // 右子序列 排序
	fmt.Println(s)
}

// 第二种方式，递归传入原数组
func quicksort2(s []int, start int, end int) {
	if start >= end {
		return
	}
	tag := s[start]
	i := start
	j := end
	for i < j {
		for s[j] > tag && i < j {
			j--
		}
		for s[i] < tag && i < j {
			i++
		}
		s[i], s[j] = s[j], s[i] //交换
	}
	quicksort2(s, start, i-1)
	quicksort2(s, i+1, end)
	fmt.Println(s)
}
