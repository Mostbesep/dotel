package commands

type CommandINFO struct {
	Section string `json:"section"` // Optional: The section to get information about, e.g., "server", "memory", "persistence"
}

type CommandCONFIGGET struct {
	Parameter string `json:"parameter"` // The configuration parameter to get
}

type CommandCONFIGSET struct {
	Parameter string `json:"parameter"` // The configuration parameter to set
	Value     string `json:"value"`     // The value to set for the configuration parameter
}

type CommandFLUSHDB struct {
	// No parameters for removing all keys from the current database
}

type CommandFLUSHALL struct {
	// No parameters for removing all keys from all databases
}

type CommandSAVE struct {
	// No parameters for synchronously saving the dataset to disk
}

type CommandBGSAVE struct {
	// No parameters for asynchronously saving the dataset to disk
}

type CommandSHUTDOWN struct {
	Save string `json:"save"` // Optional: "save" to save the dataset before shutting down, or "nosave" to not save
}
