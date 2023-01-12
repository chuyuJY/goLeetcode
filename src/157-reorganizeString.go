package main

import "container/heap"

func reorganizeString(s string) string {
	hashMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		hashMap[s[i]]++
	}
	mh := &maxHeap{}
	for val, sequence := range hashMap {
		heap.Push(mh, char{
			val:      val,
			sequence: sequence,
		})
	}
	// 构造新的字符串
	res := []byte{}
	for mh.Len() > 0 {
		char1 := heap.Pop(mh).(char)
		if mh.Len() > 0 {
			char2 := heap.Pop(mh).(char)
			res = append(res, char1.val, char2.val)
			char2.sequence--
			if char2.sequence > 0 {
				heap.Push(mh, char2)
			}
		} else {
			if len(res) > 0 && res[len(res)-1] == char1.val {
				return ""
			}
			res = append(res, char1.val)
		}
		char1.sequence--
		if char1.sequence > 0 {
			heap.Push(mh, char1)
		}
	}
	return string(res)
}

type char struct {
	val      byte
	sequence int
}

// type maxHeap []char

func (mh maxHeap) Len() int {
	return len(mh)
}

func (mh maxHeap) Less(i, j int) bool {
	if mh[i].sequence == mh[j].sequence {
		return mh[i].val < mh[j].val
	}
	return mh[i].sequence > mh[j].sequence
}

func (mh maxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *maxHeap) Push(v interface{}) {
	*mh = append(*mh, v.(char))
}

func (mh *maxHeap) Pop() interface{} {
	v := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return v
}
