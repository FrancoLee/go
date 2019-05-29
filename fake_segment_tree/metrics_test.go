package fake_segment_tree

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	m := NewMetric()
	for i := 100; i > 0; i-- {
		m.Insert(float64(i))
		time.Sleep(time.Millisecond)
	}
	time.Sleep(10*time.Second -11*time.Millisecond)
	for i := 0; i < 10; i++ {
		fmt.Println(m.Get())
		time.Sleep(time.Millisecond)
	}
	for i := 10; i > 0; i-- {
		m.Insert(float64(i))
		fmt.Println(m.Get())
		time.Sleep(time.Millisecond)
	}
	fmt.Print(m.Get())
}
