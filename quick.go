package sort

//[]int{5, 2, 12, 34, 89, 3}
func Quick(values []int) {
	if len(values) < 1 {
		return
	}

	mid, i := values[0], 1

	head, tail := 0, len(values)-1

	for head < tail {
		//fmt.Println(values)

		if values[i] > mid {
			//将大于的值放到末尾
			values[i], values[tail] = values[tail], values[i]
			tail-- //向前进
		} else {
			//将小于的放到前面
			values[i], values[head] = values[head], values[i]
			//向后比较
			head++
			i++
		}
	}
	//下一个回合
	values[head] = mid
	Quick(values[:head])   //前半部分
	Quick(values[head+1:]) //后半部分
}
