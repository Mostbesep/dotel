package commands

type CommandMULTI struct {
	// No parameters for starting a transaction
}

type CommandEXEC struct {
	// No parameters for executing all commands in a transaction
}

type CommandDISCARD struct {
	// No parameters for discarding all commands in a transaction
}

type CommandWATCH struct {
	Keys []string `json:"keys"` // List of keys to watch for changes
}
