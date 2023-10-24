package pool

import (
	"encoding/json"
	"sync"
	"unsafe"
)

type node struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	Field3 string `json:"field3"`
	Field4 string `json:"field4"`
	Field5 string `json:"field5"`
}

type graph struct {
	Nodes map[int]*node
}

type node1 struct {
	Id     int    `json:"id"`
	Name   []byte `json:"name"`
	Author []byte `json:"author"`
	Field1 []byte `json:"field1"`
	Field2 []byte `json:"field2"`
	Field3 []byte `json:"field3"`
	Field4 []byte `json:"field4"`
	Field5 []byte `json:"field5"`
}

type graph1 struct {
	Nodes map[int]*node1
}

func randChar() string {
	return "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}

func randChar1() []byte {
	str := randChar()
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func normalNew() {
	for i := 0; i < 5; i++ {
		go func() {
			graph := graph{
				Nodes: make(map[int]*node),
			}

			for j := 0; j < 1000; j++ {
				n := &node{
					Id:     j,
					Name:   randChar(),
					Author: randChar(),
					Field1: randChar(),
					Field2: randChar(),
					Field3: randChar(),
					Field4: randChar(),
					Field5: randChar(),
				}

				graph.Nodes[n.Id] = n
			}

			json.Marshal(graph)

		}()
	}
}

func poolNew() {
	pool := sync.Pool{
		New: func() any {
			return new(node)
		},
	}

	for i := 0; i < 5; i++ {
		go func() {
			graph := graph{
				Nodes: make(map[int]*node),
			}

			for j := 0; j < 1000; j++ {
				n := pool.Get().(*node)
				n.Id = j
				n.Name = randChar()
				n.Author = randChar()
				n.Field1 = randChar()
				n.Field2 = randChar()
				n.Field3 = randChar()
				n.Field4 = randChar()
				n.Field5 = randChar()

				graph.Nodes[n.Id] = n
			}
			json.Marshal(graph)

			for _, v := range graph.Nodes {
				v.Name = ""
				v.Author = ""
				v.Field1 = ""
				v.Field2 = ""
				v.Field3 = ""
				v.Field4 = ""
				v.Field5 = ""
				pool.Put(v)
			}

		}()
	}
}

func PoolNewByte() {
	pool := sync.Pool{
		New: func() any {
			return new(node1)
		},
	}

	for i := 0; i < 5; i++ {
		go func() {
			graph := graph1{
				Nodes: make(map[int]*node1),
			}

			for j := 0; j < 1000; j++ {
				n := pool.Get().(*node1)
				n.Id = j
				n.Name = randChar1()
				n.Author = randChar1()
				n.Field1 = randChar1()
				n.Field2 = randChar1()
				n.Field3 = randChar1()
				n.Field4 = randChar1()
				n.Field5 = randChar1()

				graph.Nodes[n.Id] = n
			}
			json.Marshal(graph)

			for _, v := range graph.Nodes {
				v.Name = nil
				v.Author = nil
				v.Field1 = nil
				v.Field2 = nil
				v.Field3 = nil
				v.Field4 = nil
				v.Field5 = nil
				pool.Put(v)
			}

		}()
	}
}

func Test() {
	panic("skdjf")
}
