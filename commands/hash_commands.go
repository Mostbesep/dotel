package commands

type CommandHSET struct {
	Key   string `json:"key"`   // The key of the hash
	Field string `json:"field"` // The field to set
	Value string `json:"value"` // The value to set for the field
}

type CommandHGET struct {
	Key   string `json:"key"`   // The key of the hash
	Field string `json:"field"` // The field to get
}

type CommandHGETALL struct {
	Key string `json:"key"` // The key of the hash
}

type CommandHDEL struct {
	Key    string   `json:"key"`    // The key of the hash
	Fields []string `json:"fields"` // List of fields to remove from the hash
}

type CommandHEXISTS struct {
	Key   string `json:"key"`   // The key of the hash
	Field string `json:"field"` // The field to check existence
}

type CommandHINCRBY struct {
	Key    string `json:"key"`    // The key of the hash
	Field  string `json:"field"`  // The field whose value is to be incremented
	Amount int64  `json:"amount"` // The increment amount
}

type CommandHKEYS struct {
	Key string `json:"key"` // The key of the hash
}

type CommandHVALS struct {
	Key string `json:"key"` // The key of the hash
}

type CommandHLEN struct {
	Key string `json:"key"` // The key of the hash
}

type CommandHMSET struct {
	Key    string            `json:"key"`    // The key of the hash
	Fields map[string]string `json:"fields"` // A map of field-value pairs to set in the hash
}
