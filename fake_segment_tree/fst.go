package fake_segment_tree

import (
	"math"
	"sync"
)

type Fst struct {
	max   []float64
	hsum  []int64
	sum   []float64
	mutex *sync.RWMutex
}

func NewFST() *Fst {
	fst := &Fst{
		max:   make([]float64, 2050),
		hsum:  make([]int64, 2050),
		sum:   make([]float64, 2050),
		mutex: &sync.RWMutex{},
	}
	return fst
}

func (fst *Fst) pushup(x int) {
	fst.max[x] = math.Max(fst.max[x<<1], fst.max[x<<1|1]) //更新节点(区间)最值,可以改成求和,最小
	fst.hsum[x] = fst.hsum[x<<1] + fst.hsum[x<<1|1]
	fst.sum[x] = fst.sum[x<<1] + fst.sum[x<<1|1]
}

func (fst *Fst) 	Update(index int, val float64, l int, r int, rt int) {
	fst.mutex.Lock()
	defer fst.mutex.Unlock()
	fst.doUpdate(index, val, l, r, rt)
}
func (fst *Fst) doUpdate(index int, val float64, l int, r int, rt int) {
	//[l,r]表示当前区间,rt表示当前区间编号,index表示要修改的位置,val是要改成的数值
	if l == r {
		fst.max[rt] = math.Max(val, fst.max[rt])
		fst.hsum[rt] += 1
		fst.sum[rt] += val
		return
	}
	m := (l + r) >> 1
	if index <= m { //判断目标位置在当前区间的做还是右
		fst.doUpdate(index, val, l, m, rt<<1)
	} else {
		fst.doUpdate(index, val, m+1, r, rt<<1|1)
	}
	fst.pushup(rt) //维护最值
}
func (fst *Fst) Query(L int, R int, l int, r int, rt int) (float64, int64, float64) {
	fst.mutex.RLock()
	defer fst.mutex.RUnlock()
	return fst.doQuery(L, R, l, r, rt)
}
func (fst *Fst) doQuery(L int, R int, l int, r int, rt int) (float64, int64, float64) { //L,R为要查询的区间,[l,r]当前区间
	if L <= l && r <= R { //在区间内直接返回
		return fst.max[rt], fst.hsum[rt], fst.sum[rt]
	}
	m := (l + r) >> 1
	var m1, m2, s1, s2 float64
	var h1, h2 int64
	if L <= m {
		m1, h1, s1 = fst.Query(L, R, l, m, rt<<1)
	}
	if R > m {
		m2, h2, s2 = fst.Query(L, R, m+1, r, rt<<1|1)
	}
	return math.Max(m1, m2), h1 + h2, s1 + s2
}
