type MinStack struct {
	st []int
	minSt []int
}

func Constructor() MinStack {
	st := make([]int, 0)
	minSt := make([]int, 0)
	return MinStack{
		st: st,
		minSt: minSt,
	}
}

func (this *MinStack) Push(val int) {
	this.st = append(this.st, val)
	if len(this.minSt) == 0 {
		this.minSt = append(this.minSt, val)
	} else {
		top := this.minSt[len(this.minSt)-1]
		if val < top {
			top = val
		}
		this.minSt = append(this.minSt, top)
	}
}

func (this *MinStack) Pop() {
	this.st = this.st[:len(this.st)-1]
	this.minSt = this.minSt[:len(this.minSt)-1]
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
	return this.minSt[len(this.minSt)-1]
}
