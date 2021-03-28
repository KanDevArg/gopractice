package races

import (
	"fmt"
	"sync"
)

type SafeNumber struct {
	val int
	m   sync.Mutex
}

func (i *SafeNumber) Get() int {
	i.m.Lock()
	defer i.m.Unlock()
	return i.val
}

func (i *SafeNumber) Set(val int) {
	i.m.Lock()
	defer i.m.Unlock()
	i.val = val
}

func getNumber() int {
	sfn := &SafeNumber{}
	go func() {
		sfn.Set(5)
	}()
	return sfn.Get()
}

func Races() {
	number := getNumber()
	fmt.Println(number)
}
