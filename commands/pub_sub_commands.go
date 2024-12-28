package commands

type CommandPUBLISH struct {
	Channel string `json:"channel"` // The channel to publish to
	Message string `json:"message"` // The message to publish
}

type CommandSUBSCRIBE struct {
	Channels []string `json:"channels"` // List of channels to subscribe to
}

type CommandUNSUBSCRIBE struct {
	Channels []string `json:"channels"` // List of channels to unsubscribe from
}

type CommandPSUBSCRIBE struct {
	Patterns []string `json:"patterns"` // List of patterns to subscribe to
}

type CommandPUNSUBSCRIBE struct {
	Patterns []string `json:"patterns"` // List of patterns to unsubscribe from
}
