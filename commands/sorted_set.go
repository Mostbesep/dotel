package commands

type CommandZADD struct {
	Key     string    `json:"key"`     // The key of the sorted set
	Members []ZMember `json:"members"` // List of members with their score to add
}

type ZMember struct {
	Member string  `json:"member"` // The member to add
	Score  float64 `json:"score"`  // The score of the member
}

type CommandZREM struct {
	Key     string   `json:"key"`     // The key of the sorted set
	Members []string `json:"members"` // List of members to remove
}

type CommandZSCORE struct {
	Key    string `json:"key"`    // The key of the sorted set
	Member string `json:"member"` // The member whose score is to be retrieved
}

type CommandZINCRBY struct {
	Key       string  `json:"key"`       // The key of the sorted set
	Member    string  `json:"member"`    // The member whose score is to be incremented
	Increment float64 `json:"increment"` // The amount to increment the score by
}

type CommandZRANGE struct {
	Key   string `json:"key"`   // The key of the sorted set
	Start int    `json:"start"` // The start rank for the range
	End   int    `json:"end"`   // The end rank for the range
}

type CommandZRANGEBYSCORE struct {
	Key      string  `json:"key"`       // The key of the sorted set
	MinScore float64 `json:"min_score"` // The minimum score for the range
	MaxScore float64 `json:"max_score"` // The maximum score for the range
}

type CommandZRANK struct {
	Key    string `json:"key"`    // The key of the sorted set
	Member string `json:"member"` // The member whose rank is to be retrieved
}

type CommandZCARD struct {
	Key string `json:"key"` // The key of the sorted set to get the number of members
}

type CommandZREMRANGEBYSCORE struct {
	Key      string  `json:"key"`       // The key of the sorted set
	MinScore float64 `json:"min_score"` // The minimum score for the range
	MaxScore float64 `json:"max_score"` // The maximum score for the range
}

type CommandZREVRANGE struct {
	Key   string `json:"key"`   // The key of the sorted set
	Start int    `json:"start"` // The start rank for the range (in reverse order)
	End   int    `json:"end"`   // The end rank for the range (in reverse order)
}

type CommandZREVRANK struct {
	Key    string `json:"key"`    // The key of the sorted set
	Member string `json:"member"` // The member whose rank is to be retrieved (in reverse order)
}
