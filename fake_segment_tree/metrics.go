package fake_segment_tree

import (
	"fmt"
	"math"
	"time"
)

type Metrice struct {
	fsts []*Fst
	old  int
	new  int
	time []int64
}

func NewMetric() *Metrice {
	m := &Metrice{
		old:  0,
		new:  10,
		time: make([]int64, 11),
	}
	t := time.Now().UnixNano()/1e9 - 11
	for i := 0; i < 11; i++ {
		m.fsts = append(m.fsts, NewFST())
		m.time[i] = t + int64(i) + 1
	}

	return m
}

type mhs struct {
	max  float64
	sum  float64
	hsum int64
}

func (m *Metrice) Get() (float64, float64, float64) {
	a := &mhs{}
	b := &mhs{}
	mill := time.Now().UnixNano() / 1e6
	sec := mill / 1000 //当前秒数
	fmt.Printf("%v: ", mill)
	mill %= 1000 //当前时间的毫秒偏移
	for i := 0; i < 11; i++ {
		if sec-m.time[i] == 10 {
			a.max, a.hsum, a.sum = m.fsts[i].Query(int(mill), 1000, 1, 1000, 1)
		} else if sec == m.time[i] {
			a.max, a.hsum, a.sum = m.fsts[i].Query(1, int(mill), 1, 1000, 1)
		} else if sec-m.time[i] < 10 {
			a.max, a.hsum, a.sum = m.fsts[i].Query(1, 1000, 1, 1000, 1)
		}
		calculate(a, b)
	}
	fmt.Printf("%v: ", b.hsum)
	return b.max, b.sum, b.sum / float64(b.hsum)
}
func calculate(a *mhs, b *mhs) {
	b.max = math.Max(a.max, b.max)
	b.sum = b.sum + a.sum
	b.hsum = a.hsum + b.hsum
}
func (m *Metrice) Insert(val float64) {
	begin := time.Now().UnixNano()
	now := time.Now().UnixNano()
	t := now/1e9 - m.time[m.new]
	if t > 0 {
		m.fsts[m.old] = m.fsts[m.new]
		m.time[m.old] = m.time[m.new]
		m.time[m.new] = now / 1e9
		m.old = (m.old + 1) % 10
		m.fsts[m.new] = NewFST()
	}
	m.fsts[m.new].Update(int(now/1e6)%1000, val, 1, 1000, 1)
	end := time.Now().UnixNano()
	fmt.Println(end- begin)
	//fmt.Println(int(now/1e6) % 1000)

}
