package graphs

import "fmt"

type GraphNode struct {
	neighbors map[*GraphNode]struct{}
	label     string
	color     int
}

func ColorGraph(graph map[*GraphNode]struct{}, colors []int) {

	for node := range graph {
		if _, ok := node.neighbors[node]; ok {
			fmt.Printf("legal coloring impossible for node with loop:%+v\n", node)
			return
		}

		illegalColors := map[int]struct{}{}
		for n := range node.neighbors {
			if n.color != 0 {
				// The color is used at a neighbor - we can't legally paint it here...
				illegalColors[n.color] = struct{}{}
			}
		}

		// color the node!
		for _, color := range colors {
			if _, ok := illegalColors[color]; !ok {
				node.color = color
				fmt.Printf("coloring with %+v\n", color)
				break
			}
		}
	}
}

// ShortestRoute finds the shortest route for a message from one user to another.
// It returns a list of users that makes up the route.
// You wrote a trendy new messaging app, MeshMessage, to get around flaky cell phone coverage.
//
// Instead of routing texts through cell towers, your app sends messages via the phones of nearby users, passing each message along from one phone to the next until it reaches the intended recipient. (Don't worry—the messages are encrypted while they're in transit.)
//
// Some friends have been using your service, and they're complaining that it takes a long time for messages to get delivered. After some preliminary debugging, you suspect messages might not be taking the most direct route from the sender to the recipient.
//
// Given information about active users on the network, find the shortest route for a message from one user (the sender) to another (the recipient). Return a list of users that make up this route
func ShortestRoute(graph map[*GraphNode]struct{}, startNode, endNode GraphNode) {
	//    if start_node not in graph:
	//        raise Exception('Start node not in graph')
	//    if end_node not in graph:
	//        raise Exception('End node not in graph')
	//
	//    nodes_to_visit = deque()
	//    nodes_to_visit.append(start_node)
	//
	//    # Keep track of how we got to each node
	//    # We'll use this to reconstruct the shortest path at the end
	//    # We'll ALSO use this to keep track of which nodes we've
	//    # already visited
	//    how_we_reached_nodes = {start_node: None}
	//
	//    while len(nodes_to_visit) > 0:
	//        current_node = nodes_to_visit.popleft()
	//
	//        # Stop when we reach the end node
	//        if current_node == end_node:
	//            return reconstruct_path(how_we_reached_nodes, start_node, end_node)
	//
	//        for neighbor in graph[current_node]:
	//            if neighbor not in how_we_reached_nodes:
	//                nodes_to_visit.append(neighbor)
	//                how_we_reached_nodes[neighbor] = current_node
	//
	//    # If we get here, then we never found the end node
	//    # so there's NO path from start_node to end_node
	//    return None

}

//def reconstruct_path(previous_nodes, start_node, end_node):
//    reversed_shortest_path = []
//
//    # Start from the end of the path and work backwards
//    current_node = end_node
//    while current_node:
//        reversed_shortest_path.append(current_node)
//        current_node = previous_nodes[current_node]
//
//    # Reverse our path to get the right order
//    reversed_shortest_path.reverse()  # flip it around, in place
//    return reversed_shortest_path  # no longer reversed

// 1. Use a Queue to do a BFS. Loop through nodes while len of nodes to visit is greater than 0. If a node is not in how_we_reached_nodes[neighbor], then add it to the queue and how_we_reached_nodes
// 2. Use a map to keep track of the path, e.g. how_we_reached_nodes[neighbor] = current_node
// 3. When we get to the end node, go through our map of relationships until we get to the start (recall the start node has a nil entry). Reverse that order and you have shortest path.
