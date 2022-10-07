package v1

type Context struct {
	Title  string
	Static string
}

type RequestCounts struct {
	CountObjects string `json:"countObjects,omitempty"`
	CountWindows string `json:"countWindows,omitempty"`
}

type ResponseCounts struct {
	CountObjects int `json:"countObjectsAdd,omitempty"`
	CountWindows int `json:"countWindowsAdd,omitempty"`
	ErrorObjects int `json:"errorObjects,omitempty"`
	ErrorWindows int `json:"errorWindows,omitempty"`
}
