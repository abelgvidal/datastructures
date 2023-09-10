package datastructures

type Graph struct {
   vertices int
   adjList  map[int][]int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make(map[int][]int),
	}
}

func (g *Graph) AddEdge(source, dest int) {
	g.adjList[source] = append(g.adjList[source], dest)
}

func (g *Graph) DFSUtil(vertex int, visited map[int]bool, edges []int) []int {
	visited[vertex] = true
	edges = append(edges, vertex)
	for _, v := range g.adjList[vertex] {
		if !visited[v] {
			edges = g.DFSUtil(v, visited, edges)
		}
	}
	return edges
}

func (g *Graph) DFS(startVertex int) []int{
	visited := make(map[int]bool)
	edges :=  make([]int, 0)
	edges = g.DFSUtil(startVertex, visited, edges)
	return edges
}
