/*
 * @Author: Alexleslie
 * @Date: 2022-03-12 08:55:11
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-12 09:03:44
 * @FilePath: \goCode\src\图最短路径\main.go
 * @Description:
 */

package main

import (
	hp "container/heap"
	"fmt"
)

type heap struct {
	values *minPath
}

type edge struct {
	node string
}

type graph struct {
	nodes map[string][]edge
}

type path struct {
	value int
	nodes []string
}

type minPath []path

func newGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(origin, destiny string) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny})
	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (g *graph) getPath(origin, destiny string) int {
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.value
		}

		for _, e := range g.getEdges(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				h.push(path{value: p.value + 1, nodes: append([]string{}, append(p.nodes, e.node)...)})
			}
		}
		visited[node] = true
	}

	return 0
}

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func newHeap() *heap {
	return &heap{values: &minPath{}}
}

func (h *heap) push(p path) {
	hp.Push(h.values, p)
}

func (h *heap) pop() path {
	i := hp.Pop(h.values)
	return i.(path)
}

func main() {
	// Example
	graph := newGraph()
	graph.addEdge("2", "1")
	graph.addEdge("1", "3")
	graph.addEdge("3", "4")
	graph.addEdge("1", "0")
	graph.addEdge("0", "5")
	fmt.Println(graph.getPath("2", "5"))
}
