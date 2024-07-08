package hashing_and_hash_tables

// You are given the array paths, where paths[i] = [cityAi, cityBi] means there exists a direct path going from cityAi to cityBi.
// Return the destination city, that is, the city without any path outgoing to another city.
//
// It is guaranteed that the graph of paths forms a line without any loop, therefore, there will be exactly one destination city.
//
// Example 1:
//
// Input: paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
// Output: "Sao Paulo"
// Explanation: Starting at "London" city you will reach "Sao Paulo" city which is the destination city. Your trip consist of: "London" -> "New York" -> "Lima" -> "Sao Paulo".
func destCity(paths [][]string) string {
	origins := make(map[string]struct{}, len(paths)*2)
	destinations := make(map[string]struct{}, len(paths)*2)
	for _, path := range paths {
		origins[path[0]] = struct{}{}
		destinations[path[1]] = struct{}{}
	}

	for origin := range origins {
		if _, found := destinations[origin]; found {
			delete(destinations, origin)
		}
	}

	if len(destinations) == 1 {
		for dest := range destinations {
			return dest
		}
	}
	return "error"
}
