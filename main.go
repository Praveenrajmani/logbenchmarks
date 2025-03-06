package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/klauspost/compress/zstd"
	"google.golang.org/protobuf/proto"
)

var (
	logEntries      []AuditLog
	protoLogEntries []*ProtoAuditLog
)

func init() {
	var err error
	logEntries, err = newEntries[AuditLog]("sample.log")
	if err != nil {
		panic(err)
	}
	protoLogEntries, err = newEntries[*ProtoAuditLog]("sample.log")
	if err != nil {
		panic(err)
	}
}

// Run Size Comparisons
func main() {
	logEntry := AuditLogList{
		Entries: logEntries,
	}
	protoLogEntry := &ProtoAuditLogList{
		Entries: protoLogEntries,
	}
	// Serialize logs
	jsonData, _ := json.Marshal(logEntry)
	protoData, _ := proto.Marshal(protoLogEntry)
	msgpackData, _ := logEntry.MarshalMsg(nil)
	auditLogTupleList := AuditLogTupleList{
		Entries: func() (entries []AuditLogTuple) {
			for _, logEntry := range logEntries {
				entries = append(entries, AuditLogTuple(logEntry))
			}
			return
		}(),
	}
	msgpackDataT, _ := auditLogTupleList.MarshalMsg(nil)

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

// Reads  logs from a file
func newEntries[T any](filename string) ([]T, error) {
	var logs []T

	// Open the log file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var log T
		err := json.Unmarshal(scanner.Bytes(), &log)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Skipping invalid JSON line: %s\n", scanner.Text())
			continue // Skip malformed JSON
		}
		logs = append(logs, log)
	}

	// Check for read errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return logs, nil
}
