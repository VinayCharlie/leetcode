package main

type Allocator struct {
	n        int
	memory   []int
	freeList []int
}

func Constructor(n int) Allocator {
	return Allocator{n: n, memory: make([]int, n), freeList: make([]int, n)}
}

func (this *Allocator) allocate(size int, mID int) int {
	if size > this.n {
		return -1
	}

	for i, j := 0, 0; i < this.n && j < size; i++ {
		if this.memory[i] == 0 {
			this.freeList[j] = i
			j++
		} else {
			j = 0
		}
		if j == size {
			for k := 0; k < size; k++ {
				this.memory[this.freeList[k]] = mID
			}
			return this.freeList[0]
		}
	}
	return -1
}

func (this *Allocator) free(mID int) int {
	count := 0
	for i := 0; i < this.n; i++ {
		if this.memory[i] == mID {
			this.memory[i] = 0
			count++
		}
	}
	return count
}

func main() {

}
