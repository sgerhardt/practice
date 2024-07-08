package graphs

import "testing"

func TestColorGraph(t *testing.T) {
	node0Neighbors := map[*GraphNode]struct{}{}

	node0 := &GraphNode{
		neighbors: node0Neighbors,
		label:     "node0",
		color:     0,
	}
	node1 := &GraphNode{
		neighbors: nil,
		label:     "node1",
		color:     0,
	}
	node2 := &GraphNode{
		neighbors: nil,
		label:     "node2",
		color:     0,
	}
	node3 := &GraphNode{
		neighbors: nil,
		label:     "node3",
		color:     0,
	}
	node0Neighbors[node1] = struct{}{}
	node0Neighbors[node2] = struct{}{}
	node0Neighbors[node3] = struct{}{}

	colors := []int{1, 2, 3, 4}

	n0Neighbor := map[*GraphNode]struct{}{}
	n0Neighbor[node0] = struct{}{}
	node1.neighbors = n0Neighbor
	node2.neighbors = n0Neighbor
	node3.neighbors = n0Neighbor
	node0.neighbors = node0Neighbors

	graph := map[*GraphNode]struct{}{}
	graph[node1] = struct{}{}
	graph[node2] = struct{}{}
	graph[node3] = struct{}{}
	graph[node0] = struct{}{}

	//Graph has max degree three - node0 has three neighbors
	ColorGraph(graph, colors)
}

type AdjacencyList struct {
	//List map[*GraphNode]string[]
}

func TestShortestRoute(t *testing.T) {
	//  network = {
	//    `Min`     : [`William`, `Jayden`, `Omar`],
	//    `William` : [`Min`, `Noam`],
	//    `Jayden`  : [`Min`, `Amelia`, `Ren`, `Noam`],
	//    `Ren`     : [`Jayden`, `Omar`],
	//    `Amelia`  : [`Jayden`, `Adam`, `Miguel`],
	//    `Adam`    : [`Amelia`, `Miguel`, `Sofia`, `Lucas`],
	//    `Miguel`  : [`Amelia`, `Adam`, `Liam`, `Nathan`],
	//    `Noam`    : [`Nathan`, `Jayden`, `William`],
	//    `Omar`    : [`Ren`, `Min`, `Scott`],
	//    ...
	//}
	//
	//For the network above, a message from Jayden to Adam should have this route:
	//
	//  [`Jayden`, `Amelia`, `Adam`]

	//nodes := map[*GraphNode]struct{}{}
	//Min := GraphNode{
	//	neighbors: map[*GraphNode]struct{}{[`William`, `Jayden`, `Omar`]},
	//	label:     "Min",
	//}

	//ShortestRoute(`Jayden`, `Adam`)
}
