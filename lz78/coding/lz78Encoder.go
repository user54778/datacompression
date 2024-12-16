package coding

import (
	"lz78.adpollak.net/trie"
)

type LZ78 struct {
	Trie *trie.LZ78Trie
}

type Encoded struct {
	Index  int
	Letter string
}

func (l *LZ78) LZ78Encode(sequence string) []Encoded {
	var output []Encoded
	prefix := ""
	for _, r := range sequence {
		prefix += string(r)

		if match := l.Trie.Find(prefix); match == 0 {
			// fmt.Printf("Encoder Output: %d %s\n", l.Trie.Find(prefix[:len(prefix)-1]), prefix)
			output = append(output, Encoded{
				// Index:  l.Trie.NextIndex,
				Index:  l.Trie.Find(prefix[:len(prefix)-1]),
				Letter: prefix,
			})

			l.Trie.Insert(prefix)
			prefix = ""
		}
	}
	return output
}

/*
	c := string(r)
	var index int
	if l.Trie.Find(c) == 0 {
		index = l.Trie.Insert(c)
	} else if i+1 < len(sequence) {
		concat := c + string(sequence[i+1])
		if l.Trie.Find(concat) == 0 {
			index = l.Trie.Insert(concat)
			fmt.Printf("Encoder Output: %d, %s\n", index, concat)
		}
	}
*/
// Build the longest substring in the Trie.
/*
	for i, r := range sequence {
		substring := ""
		longest := 0
		for j := i; j < len(sequence); j++ {
			substring += string(sequence[j])

			matchIndex := l.Trie.Find(substring)
			if matchIndex == 0 {
				in := l.Trie.Insert(substring)
				fmt.Printf("Encoder Output: %d, %s\n", in, string(r))
				break
			}

			// update longest matching index
			longest = matchIndex

			if j == len(sequence)-1 {
				fmt.Printf("Encoder Output: %d\n", longest)
			}
		}
		i += len(substring)
		/*
			c := string(r)
			var index int
			if l.Trie.Find(c) == 0 {
				index = l.Trie.Insert(c)
			} else if i+1 < len(sequence) {
				concat := c + string(sequence[i+1])
				if l.Trie.Find(concat) == 0 {
					index = l.Trie.Insert(concat)
					fmt.Printf("Encoder Output: %d, %s\n", index, concat)
				}
			}
*/
