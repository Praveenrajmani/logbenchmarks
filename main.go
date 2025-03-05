package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/klauspost/compress/zstd"
	"google.golang.org/protobuf/proto"
)

// Run Size Comparisons
func main() {
	logEntry := newLogEntry()
	protoLogEntry := newProtoLogEntry()

	// Serialize logs
	jsonData, _ := json.Marshal(logEntry)
	protoData, _ := proto.Marshal(protoLogEntry)
	msgpackData, _ := logEntry.MarshalMsg(nil)
	msgpackDataT, _ := AuditLogTuple(*logEntry).MarshalMsg(nil)

	// Compress logs using Zstd
	z, _ := zstd.NewWriter(nil)

	jsonDataZ := z.EncodeAll(jsonData, nil)
	protoDataZ := z.EncodeAll(protoData, nil)
	msgpackDataZ := z.EncodeAll(msgpackData, nil)
	msgpackDataTZ := z.EncodeAll(msgpackDataT, nil)

	// Print file size results
	fmt.Println("Size Comparison (Bytes):")
	fmt.Printf("JSON: %d bytes\n", len(jsonData))
	fmt.Printf("Protobuf: %d bytes\n", len(protoData))
	fmt.Printf("MessagePack: %d bytes\n", len(msgpackData))
	fmt.Printf("MessagePack Tuple: %d bytes\n", len(msgpackDataT))
	fmt.Printf("JSON + Zstd: %d bytes\n", len(jsonDataZ))
	fmt.Printf("Protobuf + Zstd: %d bytes\n", len(protoDataZ))
	fmt.Printf("MessagePack + Zstd: %d bytes\n", len(msgpackDataZ))
	fmt.Printf("MessagePack Tuple + Zstd: %d bytes\n", len(msgpackDataTZ))
}

// Create a sample log entry
func newLogEntry() *AuditLog {
	return &AuditLog{
		Timestamp: time.Now().Unix(),
		Event:     "user_login",
		User:      "admin",
	}
}

func newProtoLogEntry() *ProtoAuditLog {
	return &ProtoAuditLog{
		Timestamp: time.Now().Unix(),
		Event:     "user_login",
		User:      "admin",
	}
}
