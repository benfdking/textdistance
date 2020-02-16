package textdistance

import (
	mapset "github.com/deckarep/golang-set"
	"strings"
)

func WordsToSet(s string) mapset.Set {
	set := mapset.NewSet()
	for _, token := range strings.Fields(s) {
		set.Add(token)
	}
	return set
}

