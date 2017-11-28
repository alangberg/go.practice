package service

import (
	"sort"

	"github.com/alangberg/go.tuiter/src/domain"
)

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func containsUser(usersSlice []*domain.User, toFind *domain.User) bool {
	for _, u := range usersSlice {
		if u.Username == toFind.Username {
			return true
		}
	}
	return false
}

func containsPlugin(pluginsSlice []domain.TweetPlugin, toFind domain.TweetPlugin) bool {
	for _, p := range pluginsSlice {
		if p.GetPluginName() == toFind.GetPluginName() {
			return true
		}
	}
	return false
}
