package coding

import (
	"lz78.adpollak.net/trie"
)

func (l *LZ78) LZ78Decode(encoded []Encoded) (string, *trie.LZ78Trie) {
	decoded := ""
	t := trie.NewLZ78Trie()

	for _, token := range encoded {
		var reconstruct string
		if token.Index == 0 {
			reconstruct = token.Letter
			t.Insert(reconstruct)
		} else {
			t.Insert(token.Letter)
			reconstruct = token.Letter
		}

		decoded += reconstruct
	}
	return decoded, t
}
