package unbounded

import (
	"math"
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
	Symbols    map[string]Symbol
	TotalCount int
}

func NewModel(totalCount int) *Model {
	return &Model{
		Symbols:    make(map[string]Symbol),
		TotalCount: totalCount,
	}
}

func (m *Model) ComputeCount(symbols []rune) {
	for _, r := range symbols {
		if v, ok := m.Symbols[string(r)]; ok {
			v.Count++
			m.Symbols[string(r)] = v
		} else {
			m.Symbols[string(r)] = Symbol{Count: 1}
		}
		m.TotalCount++
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

	cumProb := 0.0
	for _, kv := range ss {
		p := float64(kv.Value.Count) / float64(m.TotalCount)
		s := NewSymbol(p, cumProb, kv.Value.Count)
		cumProb += p
		m.Symbols[kv.Key] = *s
		// fmt.Printf("Symbol: %c, Probability: %f, CumulativeProbability: %f\n", kv.Key, s.CurrentProbability, s.CumulativeProbability)
	}
}

type kv struct {
	Key   string
	Value Symbol
}

func roundProbability(f float64) float64 {
	return math.Round(f*100) / 100
}
