package main

import "fmt"


func bubbleSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := 0; j < len(a)-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func quickSort(a []int, l, r int) {
	if l < r {
		p := func() int {
			v := a[(l+r)/2]
			i := l
			j := r
			for i <= j {
				for a[i] < v {
					i++
				}
				for a[j] > v {
					j--
				}
				if i >= j {
					break
				}
				a[i], a[j] = a[j], a[i]
				i++
				j--
			}
			return j
		}()
		quickSort(a, l, p)
		quickSort(a, p+1, r)
	}
}

func mergeSort(a []int) []int {
	count := len(a)

	if count > 2 {
		lb := mergeSort(a[:count/2])
		rb := mergeSort(a[count/2:])
		a := append(lb, rb...)
		lastIndex := len(a) - 1

		for i := 0; i < lastIndex; i++ {
			mv := a[i]
			mi := i

			for j := i + 1; j < lastIndex+1; j++ {
				if mv > a[j] {
					mv = a[j]
					mi = j
				}
			}
			if mi!=i{
				a[i], a[mi]=a[mi], a[i]
			}
		}
	}else if len(a)>1 && a[0]>a[1]{
		a[0], a[1]=a[1], a[0]
	}
	return a
}

func main() {
	a := []int{1, 2, 9, 8, 4, 3}
	fmt.Println(a)
	// bubbleSort(a)
	// quickSort(a, 0, len(a)-1)
	mergeSort(a)
	fmt.Println(a)
}
