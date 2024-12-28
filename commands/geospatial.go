package commands

type CommandGEOADD struct {
	Key       string  `json:"key"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Member    string  `json:"member"`
}

type CommandGEODIST struct {
	Key     string `json:"key"`
	Member1 string `json:"member1"`
	Member2 string `json:"member2"`
}

type CommandGEORADIUS struct {
	Key       string  `json:"key"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Radius    float64 `json:"radius"`
	Unit      string  `json:"unit"` // e.g., km, m, mi, ft
}

type CommandGEORADIUSBYMEMBER struct {
	Key    string  `json:"key"`
	Member string  `json:"member"`
	Radius float64 `json:"radius"`
	Unit   string  `json:"unit"` // e.g., km, m, mi, ft
}
