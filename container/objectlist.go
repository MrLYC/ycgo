package container

import (
	"container/list"
)

type ObjectList struct {
	list.List
}

func (l *ObjectList) Each(f func(int, interface{}) error) (int, error) {
	var (
		i   = 0
		err error
	)

	for e := l.Front(); e != nil; e = e.Next() {
		err = f(i, e.Value)
		if err != nil {
			break
		}
		i = i + 1
	}
	return i, err
}
