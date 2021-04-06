package datastructure

// Size of node
const Size = 26

// Node represent each node in tries
type Node struct {
	Children [Size]*Node
	IsEnd    bool
}

// Tries represent a tries and has pointer to node
type Tries struct {
	Root *Node
}

// InitTries to init tries data structure
func InitTries() *Tries {
	return &Tries{
		Root: &Node{},
	}
}

// Insert to insert string to tries data structure
func (t *Tries) Insert(s string) {
	wl := len(s)
	currentNode := t.Root
	for i := 0; i < wl; i++ {
		// becuase a is 97 ASCII but we need a to start from 0 because children contain array of node size 26 alphabet charactor
		charIndex := s[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			currentNode.Children[charIndex] = &Node{}
		}
		// to go to next node
		currentNode = currentNode.Children[charIndex]
	}
	currentNode.IsEnd = true
}

// Search
func (t *Tries) Search(s string) bool {
	wl := len(s)
	currentNode := t.Root
	for i := 0; i < wl; i++ {
		// becuase a is 97 ASCII but we need a to start from 0 because children contain array of node size 26 alphabet charactor
		charIndex := s[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			return false
		}
		// to go to next node
		currentNode = currentNode.Children[charIndex]
	}

	if currentNode.IsEnd {
		return currentNode.IsEnd
	}

	return false
}
