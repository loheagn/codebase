package leetcode

type Node struct {
	val       int
	zero, one *Node
}

func maximizeXor(nums []int, queries [][]int) []int {
	root := &Node{}

	add := func(x int) {
		p := root
		for i := 31; i >= 0; i-- {
			t := (x >> i) & 1
			if t == 0 {
				if p.zero == nil {
					p.zero = &Node{val: 0}
				}
				p = p.zero
			} else {
				if p.one == nil {
					p.one = &Node{val: 1}
				}
				p = p.one
			}
		}
	}

	for _, v := range nums {
		add(v)
	}

	search := func(x, m int) int {
		p := root
		mOK := false
		re := 0
		for i := 31; i >= 0; i-- {
			t := (x >> i) & 1
			var next *Node
			if t^1 > t {
				if p.one != nil {
					next = p.one
				} else {
					next = p.zero
				}
			} else {
				if p.zero != nil {
					next = p.zero
				} else {
					next = p.one
				}
			}
			if !mOK {
				tm := (m >> i) & 1
				if tm > next.val {
					if p.zero == nil {
						return -1
					}
					next = p.zero
				}
				if next.val < tm {
					mOK = true
				}
			}
			if next == nil {
				return -1
			}
			p, re = next, re^(next.val<<i)
		}
		return re
	}

	re := make([]int, 0, len(queries))
	for _, arr := range queries {
		re = append(re, search(arr[0], arr[1]))
	}

	return re
}
