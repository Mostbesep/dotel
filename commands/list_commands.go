package commands

type CommandLPUSH struct {
	Key    string   `json:"key"`    // The key to prepend to
	Values []string `json:"values"` // List of values to prepend to the key
}

type CommandRPUSH struct {
	Key    string   `json:"key"`    // The key to append to
	Values []string `json:"values"` // List of values to append to the key
}

type CommandLPOP struct {
	Key string `json:"key"` // The key to remove and get the first element from
}

type CommandRPOP struct {
	Key string `json:"key"` // The key to remove and get the last element from
}

type CommandLRANGE struct {
	Key   string `json:"key"`   // The key to get the range from
	Start int    `json:"start"` // The start index of the range
	End   int    `json:"end"`   // The end index of the range
}

type CommandLLEN struct {
	Key string `json:"key"` // The key to get the length of the list
}

type CommandLREM struct {
	Key   string `json:"key"`   // The key to remove elements from
	Count int    `json:"count"` // Number of elements to remove (0 for all occurrences)
	Value string `json:"value"` // The value to remove
}

type CommandLSET struct {
	Key   string `json:"key"`   // The key containing the list
	Index int    `json:"index"` // The index of the element to set
	Value string `json:"value"` // The new value for the element
}

type CommandLINSERT struct {
	Key      string `json:"key"`       // The key containing the list
	Before   bool   `json:"before"`    // Whether to insert before or after the reference element
	RefValue string `json:"ref_value"` // The reference element to insert before or after
	NewValue string `json:"new_value"` // The new value to insert
}
