package status

type Status struct {
	Index  string `json:"_index"`
	Id     string `json:"_id"`
	Result string `json:"result"`
	Found  bool   `json:"found"`
}
