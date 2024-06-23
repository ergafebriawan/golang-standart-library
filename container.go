package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	//container list
	data := list.New()
	data.PushBack("erga")
	data.PushBack("febriawan")
	data.PushBack("febrian")

	for e := data.Front(); e!= nil; e = e.Next() {
        println(e.Value.(string))
    }

	//container list
	data2 := ring.New(5)
	for i := 0; i < data2.Len(); i++ {
		data2.Value = "value "+ strconv.FormatInt(int64(i), 10)
		data2 = data2.Next()
	}

	data2.Do(func (value interface{})  {
		fmt.Println(value)
	})
}