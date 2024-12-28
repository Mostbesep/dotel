package commands

type CommandSETBIT struct {
	Key    string `json:"key"`
	Offset int    `json:"offset"`
	Value  int    `json:"value"` // 0 or 1
}

type CommandGETBIT struct {
	Key    string `json:"key"`
	Offset int    `json:"offset"`
}

type CommandBITCOUNT struct {
	Key string `json:"key"`
}

type CommandBITOP struct {
	Operation   string   `json:"operation"` // AND, OR, XOR, NOT
	Destination string   `json:"destination"`
	Sources     []string `json:"sources"`
}
