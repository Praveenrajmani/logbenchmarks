package main

//go:generate msgp -tests=false $GOFILE

// Import generated Protobuf file
type AuditLog struct {
	Timestamp int64  `json:"timestamp"`
	Event     string `json:"event"`
	User      string `json:"user"`
}

//msgp:tuple AuditLogTuple
type AuditLogTuple struct {
	Timestamp int64  `json:"timestamp"`
	Event     string `json:"event"`
	User      string `json:"user"`
}
