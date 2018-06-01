package search

import (
	itemDomain "github.com/audio35444/tp2-api/src/app/domain/item"
)

type Search struct {
	Hits struct {
		Total int `json:"total"`
		Hits  []struct {
			Index  string    `json:"_index"`
			Type   string    `json:"_type"`
			Id     string    `json:"_id"`
			Score  float32   `json:"_score"`
			Source itemDomain.Item `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
