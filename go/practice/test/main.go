package main

import (
	"encoding/json"
	"fmt"
)

type Range [2]int

type Node struct {
	Name     string
	NodeList []*Node
}

func (n *Node) Marshal() []byte {
	m := n.String()

	b, _ := json.Marshal(m)
	return b
}

func (n *Node) Unmarshal(b []byte) {
	m := make(map[string]any)

	json.Unmarshal(b, &m)

	b, _ = json.Marshal(m)
	fmt.Println(string(b))

	n.rec(m)
}

func (n *Node) rec(km any) {
	m := km.(map[string]any)
	// only one
	for k, v := range m {
		n.Name = k

		if vv, ok := v.([]any); ok {
			for _, vvv := range vv {
				if vvvv, ok := vvv.(string); ok {
					n.NodeList = append(n.NodeList, &Node{Name: vvvv})
				} else {
					aNode := new(Node)
					n.NodeList = append(n.NodeList, aNode)
					aNode.rec(vvv)
				}
			}
		}
	}
}

func (n *Node) String() any {
	for {
	a:

		for {
			goto a
		}
	}
}
