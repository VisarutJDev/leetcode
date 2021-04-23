package algorithm

// DFS is breadth first search function
func (b *Tree) DFS() []int {
	stack := []*Tree{}
	stack = append(stack, b)
	result := []int{b.Value}
	visited := map[int]bool{
		b.Value: true,
	}
	return DFSUtil(stack, result, visited)
}

//DFSUtil, This appends the value of the root of subtree or tree to the
//result
func DFSUtil(stack []*Tree, res []int, visited map[int]bool) []int {
	if len(stack) == 0 {
		return res
	}

	subTree := stack[len(stack)-1] //acessing last element
	switch {
	case subTree.Left != nil && visited[subTree.Left.Value] != true:

		res = append(res, subTree.Left.Value)
		visited[subTree.Left.Value] = true
		stack = append(stack, subTree.Left)
		return DFSUtil(stack, res, visited)

	case subTree.Right != nil && visited[subTree.Right.Value] != true:
		res = append(res, subTree.Right.Value)
		visited[subTree.Right.Value] = true
		stack = append(stack, subTree.Right)
		return DFSUtil(stack, res, visited)

	default:
		stack = stack[:len(stack)-1]
		return DFSUtil(stack, res, visited)

	}
}

func (b *Tree) RouteTwoNode() bool {
	stack := []*Tree{}
	stack = append(stack, b)

	visited := map[int]bool{
		b.Value: true,
	}
	return search(stack, visited, b, b.Right)
}

func search(stack []*Tree, visited map[int]bool, start, end *Tree) bool {
	if len(stack) == 0 {
		return false
	}

	subTree := stack[len(stack)-1] //acessing last element
	switch {
	case subTree.Left != nil && visited[subTree.Left.Value] != true:

		visited[subTree.Left.Value] = true
		stack = append(stack, subTree.Left)
		if subTree.Left == end {
			return true
		}
		return search(stack, visited, start, end)

	case subTree.Right != nil && visited[subTree.Right.Value] != true:
		visited[subTree.Right.Value] = true
		stack = append(stack, subTree.Right)
		if subTree.Right == end {
			return true
		}
		return search(stack, visited, start, end)

	default:
		stack = stack[:len(stack)-1]
		return search(stack, visited, start, end)

	}
}
