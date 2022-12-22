package main

import (
	"fmt"
	"reflect"
)

type Container interface {
	begin() Iterator
	append(v interface{})
}

type Iterator interface {
	Next() Iterator
	isEnd() bool
	Deref() interface{}
}

type List struct {
  head *Node
  tail *Node
  iter *Node
}

type Node struct {
  data interface{}
  next *Node
}

func (x * List) append(v interface{}) {
	tail := Node{v, nil}
	x.tail.next = &tail
	x.tail = &tail
}

func (v * List) begin() Iterator {
	v.iter = v.head
  	return &List{
        iter: v.iter,
   }
}

func (v * List) Next() Iterator {
	v.iter = v.iter.next
	if (v.isEnd()) {
		return nil
	}
	return &List{
        iter: v.iter,
    }
}

func (v List) isEnd() bool {
	return v.iter == v.tail
}

func (v List) Deref() interface{} {
	return v.iter.data
}


type Vector struct {
	arr []interface{}
	iter interface{}
	index int
	size int
}

func (x * Vector) append(v interface{}) {
	x.arr = append(x.arr,v)
	x.size++
}

func (v * Vector) begin() Iterator {
	v.iter = v.arr[0]
  	return &Vector{
        iter: v.iter,
		index: 0,
		size: v.size,
		arr: v.arr,
   }
}

func (v * Vector) Next() Iterator {
	v.index++
	if (v.isEnd()) {
		return nil
	}
	v.iter = v.arr[v.index]
	return &Vector{
        iter: v.iter,
		index: v.index,
		size: v.size,
		arr: v.arr,
    }
}


func (v Vector) isEnd() bool {
	return v.size == v.index
}

func (v Vector) Deref() interface{} {
	return v.iter
}

func SumInt(c Container) int {
  sum := 0
  for i := c.begin(); !i.isEnd(); i.Next() {
    if(reflect.TypeOf(i.Deref()).Kind() == reflect.Int) {
        sum += i.Deref().(int)
    }
  }
  return sum
}

func SumFloat64(c Container) float64 {
  sum := 0.0

  for i := c.begin(); !i.isEnd(); i.Next() {
    if(reflect.TypeOf(i.Deref()).Kind() == reflect.Float64) {
        sum += i.Deref().(float64)
    }
  }
  return sum
}

func main() {
	x := Node{50.5,nil}
	v := List{&x,&x, nil}
	fmt.Println("Appending To List....")
	v.append(10)
	v.append(11.1)
	v.append(11)
	v.append(15)
	fmt.Println("TESTING LIST:")
	for i := v.begin(); !i.isEnd(); i.Next() {
    	fmt.Println(i.Deref())
    }

	fmt.Println()

	m := Vector{nil, nil, 0, 0}
	fmt.Println("Appending To Vector....")
	m.append(10.66)
	m.append(11)
	m.append(12)
	m.append(13.25)
	m.append(14)
	fmt.Println("TESTING VECTOR:")
	for i := m.begin(); !i.isEnd(); i.Next() {
    	fmt.Println(i.Deref())
    }

	fmt.Println()
	
	p := SumInt(&m)
	fmt.Println("Adding Ints from Vector:", p)
	q := SumInt(&v)
	fmt.Println("Adding Ints from List:", q)
	
	fmt.Println()

	w := SumFloat64(&m)
	fmt.Println("Adding Float64 from Vector:", w)
	e := SumFloat64(&v)
	fmt.Println("Adding Float64 from Vector:", e)
}


