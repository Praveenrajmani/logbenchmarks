package main

// Import generated Protobuf file
import (
	"time"
)

//go:generate msgp -tests=false $GOFILE

// Grab key names from json tags..
//msgp:tag json

// ObjectVersion object version key/versionId
type ObjectVersion struct {
	ObjectName string `json:"objectName"`
	VersionID  string `json:"versionId,omitempty"`
}

// AuditLogList represents the list of audit entries
type AuditLogList struct {
	Entries []AuditLog `json:"entries"`
}

// AuditLog - audit entry logs.
type AuditLog struct {
	Version      string    `json:"version"`
	DeploymentID string    `json:"deploymentid,omitempty"`
	Time         time.Time `json:"time"`
	Event        string    `json:"event"`

	// Class of audit message - S3, admin ops, bucket management
	Type string `json:"type,omitempty"`

	// deprecated replaced by 'Event', kept here for some
	// time for backward compatibility with k8s Operator.
	Trigger string `json:"trigger"`
	API     struct {
		Name                string          `json:"name,omitempty"`
		Bucket              string          `json:"bucket,omitempty"`
		Object              string          `json:"object,omitempty"`
		Objects             []ObjectVersion `json:"objects,omitempty"`
		Status              string          `json:"status,omitempty"`
		StatusCode          int             `json:"statusCode,omitempty"`
		InputBytes          int64           `json:"rx"`
		OutputBytes         int64           `json:"tx"`
		HeaderBytes         int64           `json:"txHeaders,omitempty"`
		TimeToFirstByte     string          `json:"timeToFirstByte,omitempty"`
		TimeToFirstByteInNS string          `json:"timeToFirstByteInNS,omitempty"`
		TimeToResponse      string          `json:"timeToResponse,omitempty"`
		TimeToResponseInNS  string          `json:"timeToResponseInNS,omitempty"`
	} `json:"api"`
	RemoteHost string                 `json:"remotehost,omitempty"`
	RequestID  string                 `json:"requestID,omitempty"`
	UserAgent  string                 `json:"userAgent,omitempty"`
	ReqPath    string                 `json:"requestPath,omitempty"`
	ReqHost    string                 `json:"requestHost,omitempty"`
	ReqClaims  map[string]interface{} `json:"requestClaims,omitempty"`
	ReqQuery   map[string]string      `json:"requestQuery,omitempty"`
	ReqHeader  map[string]string      `json:"requestHeader,omitempty"`
	RespHeader map[string]string      `json:"responseHeader,omitempty"`
	Tags       map[string]interface{} `json:"tags,omitempty"`

	AccessKey  string `json:"accessKey,omitempty"`
	ParentUser string `json:"parentUser,omitempty"`

	Error string `json:"error,omitempty"`
}

// AuditLogTupleList represents the list of audit log tuples
type AuditLogTupleList struct {
	Entries []AuditLogTuple `json:"entries"`
}

//msgp:tuple AuditLogTuple
type AuditLogTuple struct {
	Version      string    `json:"version"`
	DeploymentID string    `json:"deploymentid,omitempty"`
	Time         time.Time `json:"time"`
	Event        string    `json:"event"`

	// Class of audit message - S3, admin ops, bucket management
	Type string `json:"type,omitempty"`

	// deprecated replaced by 'Event', kept here for some
	// time for backward compatibility with k8s Operator.
	Trigger string `json:"trigger"`
	API     struct {
		Name                string          `json:"name,omitempty"`
		Bucket              string          `json:"bucket,omitempty"`
		Object              string          `json:"object,omitempty"`
		Objects             []ObjectVersion `json:"objects,omitempty"`
		Status              string          `json:"status,omitempty"`
		StatusCode          int             `json:"statusCode,omitempty"`
		InputBytes          int64           `json:"rx"`
		OutputBytes         int64           `json:"tx"`
		HeaderBytes         int64           `json:"txHeaders,omitempty"`
		TimeToFirstByte     string          `json:"timeToFirstByte,omitempty"`
		TimeToFirstByteInNS string          `json:"timeToFirstByteInNS,omitempty"`
		TimeToResponse      string          `json:"timeToResponse,omitempty"`
		TimeToResponseInNS  string          `json:"timeToResponseInNS,omitempty"`
	} `json:"api"`
	RemoteHost string                 `json:"remotehost,omitempty"`
	RequestID  string                 `json:"requestID,omitempty"`
	UserAgent  string                 `json:"userAgent,omitempty"`
	ReqPath    string                 `json:"requestPath,omitempty"`
	ReqHost    string                 `json:"requestHost,omitempty"`
	ReqClaims  map[string]interface{} `json:"requestClaims,omitempty"`
	ReqQuery   map[string]string      `json:"requestQuery,omitempty"`
	ReqHeader  map[string]string      `json:"requestHeader,omitempty"`
	RespHeader map[string]string      `json:"responseHeader,omitempty"`
	Tags       map[string]interface{} `json:"tags,omitempty"`

	AccessKey  string `json:"accessKey,omitempty"`
	ParentUser string `json:"parentUser,omitempty"`

	Error string `json:"error,omitempty"`
}
