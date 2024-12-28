package commands

type CommandEVAL struct {
	Script    string   `json:"script"`    // The Lua script to execute
	Keys      []string `json:"keys"`      // The keys to pass to the Lua script
	Arguments []string `json:"arguments"` // The arguments to pass to the Lua script
}

type CommandEVALSHA struct {
	ScriptSHA string   `json:"scriptSHA"` // The SHA1 hash of the Lua script
	Keys      []string `json:"keys"`      // The keys to pass to the Lua script
	Arguments []string `json:"arguments"` // The arguments to pass to the Lua script
}

type CommandSCRIPTLOAD struct {
	Script string `json:"script"` // The Lua script to load into Redis
}

type CommandSCRIPTCHECK struct {
	ScriptSHA string `json:"scriptSHA"` // The SHA1 hash of the Lua script to check
}

type CommandSCRIPTFLUSH struct {
	// No parameters for flushing all Lua scripts
}
