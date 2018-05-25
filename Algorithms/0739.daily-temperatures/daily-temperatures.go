package problem0739

func dailyTemperatures(temperatures []int) []int {
	lens := len(temperatures)
	res := make([]int, lens)


	// 删除栈顶元素

	stack := make([]int, 0)
	top := -1
	for i := 0; i < lens; i++ {
		for top >= 0 && temperatures[i] > temperatures[stack[top]] {
			res[stack[top]] = i - stack[top]

			// 除去栈顶元素
			stack =  append(stack[:top], stack[top+1:]...)
			top--
		}

		top++
		stack[top] = i
	}


	return res
}
