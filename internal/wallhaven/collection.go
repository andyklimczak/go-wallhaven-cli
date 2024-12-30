package wallhaven

import "fmt"

type Collection struct {
	Id     int    `json:"id"`
	Label  string `json:"label"`
	Views  int    `json:"views"`
	Public int    `json:"public"`
	Count  int    `json:"count"`
}

type CollectionData struct {
	Data []Collection `json:"data"`
}

func (cd *CollectionData) GetByLabel(label string) (*Collection, error) {
	for _, c := range cd.Data {
		if c.Label == label {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("Unable to get collection by label %s", label)
}
