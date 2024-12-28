package commands

type CommandPFADD struct {
	Key      string   `json:"key"`
	Elements []string `json:"elements"`
}

type CommandPFCOUNT struct {
	Key string `json:"key"`
}

type CommandPFMERGE struct {
	Destination string   `json:"destination"`
	SourceKeys  []string `json:"source_keys"`
}
