package commands

type CommandSET struct {
	Key   string `json:"key"`   // The key to set
	Value string `json:"value"` // The value to set for the key
}

type CommandGET struct {
	Key string `json:"key"` // The key to get the value for
}

type CommandAPPEND struct {
	Key   string `json:"key"`   // The key to append to
	Value string `json:"value"` // The value to append to the key
}

type CommandINCR struct {
	Key string `json:"key"` // The key to increment
}

type CommandDECR struct {
	Key string `json:"key"` // The key to decrement
}

type CommandMGET struct {
	Keys []string `json:"keys"` // List of keys to get values for
}

type CommandMSET struct {
	KeysValues map[string]string `json:"keys_values"` // Map of key-value pairs to set
}

type CommandSETNX struct {
	Key   string `json:"key"`   // The key to set if not exists
	Value string `json:"value"` // The value to set for the key
}

type CommandSETEX struct {
	Key     string `json:"key"`     // The key to set
	Value   string `json:"value"`   // The value to set for the key
	Seconds int    `json:"seconds"` // Expiration time in seconds
}

type CommandPSETEX struct {
	Key          string `json:"key"`          // The key to set
	Value        string `json:"value"`        // The value to set for the key
	Milliseconds int    `json:"milliseconds"` // Expiration time in milliseconds
}
