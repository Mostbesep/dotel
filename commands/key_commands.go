package commands

type CommandDEL struct {
	Keys []string `json:"keys"` // List of keys to delete
}

type CommandEXISTS struct {
	Key string `json:"key"` // Key to check for existence
}

type CommandEXPIRE struct {
	Key     string `json:"key"`     // Key to set expiration on
	Seconds int    `json:"seconds"` // Time in seconds for expiration
}

type CommandTTL struct {
	Key string `json:"key"` // Key to get the TTL for
}

type CommandPERSIST struct {
	Key string `json:"key"` // Key to remove the expiration from
}

type CommandTYPE struct {
	Key string `json:"key"` // Key to get the type of
}

type CommandRENAME struct {
	OldKey string `json:"old_key"` // The existing key to rename
	NewKey string `json:"new_key"` // The new key name
}

type CommandRENAMENX struct {
	OldKey string `json:"old_key"` // The existing key to rename
	NewKey string `json:"new_key"` // The new key name, only renamed if new key does not exist
}

type CommandKEYS struct {
	Pattern string `json:"pattern"` // Pattern to match keys
}

type CommandSCAN struct {
	Cursor  int    `json:"cursor"`  // Cursor to continue scanning
	Pattern string `json:"pattern"` // Pattern to match keys
	Count   int    `json:"count"`   // Number of keys to return per call
}
