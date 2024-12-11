package unbounded

import (
	"fmt"
	"sort"
)

type Symbol struct {
	CurrentProbability    float64
	CumulativeProbability float64
	Count                 int
}

func NewSymbol(currentProbability, cumulativeProbability float64, count int) *Symbol {
	return &Symbol{
		CurrentProbability:    currentProbability,
		CumulativeProbability: cumulativeProbability,
		Count:                 count,
	}
}

type Model struct {
	Symbols    map[rune]Symbol
	TotalCount int
}

func NewModel(totalCount int) *Model {
	return &Model{
		Symbols:    make(map[rune]Symbol),
		TotalCount: totalCount,
	}
}

func (m *Model) ComputeCount(symbols []rune) {
	for _, r := range symbols {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			if v, ok := m.Symbols[r]; ok {
				v.Count++
				m.Symbols[r] = v
			} else {
				m.Symbols[r] = Symbol{Count: 1}
			}
			m.TotalCount++
		}
	}
}

func (m *Model) ComputeCumulative() {
	var ss []kv
	for k, v := range m.Symbols {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value.Count > ss[j].Value.Count
	})

	cum := 0.0
	for _, kv := range ss {
		p := float64(kv.Value.Count) / float64(m.TotalCount)
		cum += p
		s := NewSymbol(p, cum, kv.Value.Count)
		m.Symbols[kv.Key] = *s
		fmt.Printf("Symbol: %c, Probability: %f, CumulativeProbability: %f\n", kv.Key, s.CurrentProbability, s.CumulativeProbability)
	}
}

type kv struct {
	Key   rune
	Value Symbol
}
