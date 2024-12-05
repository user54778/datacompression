package tunstalltree

// IMPLEMENT ME: Implement an n-bit Tunstall code
// for a source that generates iid letters from an alphabet of size N.
// This should be done with a Heap.

// Step 1: Define the node structure of the tree
type Node struct {
	Sequence    string
	Probability float64
}

// Step 2: We can use a Node slice to implement on the Heap interface.
// Step 3: Implement Len, Less, Swap, Push, and Pop to satisify it (make sure to make it a MAX-heap).
type MaxHeap []Node

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Probability > h[j].Probability }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and pop use pointer receivers since they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Node))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Step 4: Implement the Tunstall Coding using this heap.
// Step 5: Generate a codebook from the leaves of this.
func BuildTunstallCode(maxHeap MaxHeap, n int) map[string]string {
	return nil
}

// Step 6: Create an encoder, which will simply match the longest sequence, and append the corresponding
//
//	binary code for the matched sequence to the output.
func TunstallEncode() string {
	return ""
}

// Step 7: Create a decoder, which will reverse this process.
func TunstallDecode() string {
	return ""
}
