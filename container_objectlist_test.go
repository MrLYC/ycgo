package ycgo

import (
	"testing"
	"ycgo/container"
)

func TestObjectListEach(t *testing.T) {
	list := &container.ObjectList{}
	list.PushBack(1)
	list.PushBack(2)
	arr := make([]int, list.Len())
	list.Each(func(i int, e interface{}) error {
		arr[i] = e.(int)
		return nil
	})
	if len(arr) != 2 || arr[0] != 1 || arr[1] != 2 {
		t.Errorf("Each: %v", arr)
	}
}
