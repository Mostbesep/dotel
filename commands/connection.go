package commands

type CommandAUTH struct {
	Password string `json:"password"` // The password to authenticate to the server
}

type CommandPING struct {
	// No parameters for pinging the server
}

type CommandQUIT struct {
	// No parameters for closing the connection
}
