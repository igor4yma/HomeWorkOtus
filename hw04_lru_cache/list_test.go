package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)          // [10]
		l.PushBack(20)           // [10, 20]
		l.PushBack(30)           // [10, 20, 30]
		middle := l.Front().next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		require.Equal(t, 70, l.Front().Value)
		require.Equal(t, 50, l.Back().Value)

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil && i.Value != nil; i = i.next {
			//k, ok := i.Value.(int)
			//fmt.Println(k, ok)
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

	t.Run("additional complex", func(t *testing.T) {
		l := NewList()

		l.Remove(l.Back()) // []
		require.Equal(t, 0, l.Len())

		l.Remove(l.Front()) // []
		require.Equal(t, 0, l.Len())

		l.PushBack(nil)    // [nil]
		l.PushFront(10)    // [10, nil]
		l.PushFront(20.56) // [20.56, 10, nil]
		l.PushBack('\n')   // [20.56, 10, nil, '\n']
		l.PushBack("mumu") // [20.56, 10, nil, '\n', "mumu"]
		require.Equal(t, 5, l.Len())

		l.Remove(nil)
		require.Equal(t, 5, l.Len())
	})
}
