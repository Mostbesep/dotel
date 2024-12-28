package commands

type CommandSADD struct {
	Key     string   `json:"key"`     // The key to which members are to be added
	Members []string `json:"members"` // List of members to add to the set
}

type CommandSREM struct {
	Key     string   `json:"key"`     // The key from which members are to be removed
	Members []string `json:"members"` // List of members to remove from the set
}

type CommandSMEMBERS struct {
	Key string `json:"key"` // The key to get all members from
}

type CommandSISMEMBER struct {
	Key    string `json:"key"`    // The key to check for the member
	Member string `json:"member"` // The member to check for existence
}

type CommandSCARD struct {
	Key string `json:"key"` // The key to get the number of members of
}

type CommandSMOVE struct {
	SourceKey string `json:"source_key"` // The source set key
	DestKey   string `json:"dest_key"`   // The destination set key
	Member    string `json:"member"`     // The member to move
}

type CommandSINTER struct {
	Sets []string `json:"sets"` // List of keys for which to get the intersection
}

type CommandSUNION struct {
	Sets []string `json:"sets"` // List of keys for which to get the union
}

type CommandSDIFF struct {
	Sets []string `json:"sets"` // List of keys for which to get the difference
}
