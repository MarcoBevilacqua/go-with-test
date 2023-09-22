package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type league []Player

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("problem parsing league %v", err)
	}

	return league, err
}

func (l league) find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}

	return nil
}