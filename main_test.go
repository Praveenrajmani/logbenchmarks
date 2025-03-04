package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/zstd"
	"github.com/vmihailenco/msgpack/v5"

	"google.golang.org/protobuf/proto"
)

// Import generated Protobuf file
type AuditLog struct {
	Timestamp int64  `json:"timestamp"`
	Event     string `json:"event"`
	User      string `json:"user"`
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

// Benchmark JSON Serialization
func BenchmarkJSONSerialization(b *testing.B) {
	logEntry := newLogEntry()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(logEntry)
	}
}

// Benchmark JSON Deserialization
func BenchmarkJSONDeserialization(b *testing.B) {
	logEntry := newLogEntry()
	data, _ := json.Marshal(logEntry)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var log AuditLog
		_ = json.Unmarshal(data, &log)
	}
}

// Benchmark JSON + Zstd Compression
func BenchmarkJSONZstdCompression(b *testing.B) {
	logEntry := newLogEntry()
	data, _ := json.Marshal(logEntry)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		writer, _ := zstd.NewWriter(&buf)
		_, _ = writer.Write(data)
		writer.Close()
	}
}

// Benchmark Protobuf Serialization
func BenchmarkProtobufSerialization(b *testing.B) {
	logEntry := newProtoLogEntry()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(logEntry)
	}
}

// Benchmark Protobuf Deserialization
func BenchmarkProtobufDeserialization(b *testing.B) {
	logEntry := newProtoLogEntry()
	data, _ := proto.Marshal(logEntry)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var log ProtoAuditLog
		_ = proto.Unmarshal(data, &log)
	}
}

// Benchmark Protobuf + Zstd Compression
func BenchmarkProtobufZstdCompression(b *testing.B) {
	logEntry := newProtoLogEntry()
	data, _ := proto.Marshal(logEntry)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		writer, _ := zstd.NewWriter(&buf)
		_, _ = writer.Write(data)
		writer.Close()
	}
}

// Benchmark MessagePack Serialization
func BenchmarkMessagePackSerialization(b *testing.B) {
	logEntry := newLogEntry()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(logEntry)
	}
}

// Benchmark MessagePack + Zstd Compression
func BenchmarkMessagePackZstdCompression(b *testing.B) {
	logEntry := newLogEntry()
	data, _ := msgpack.Marshal(logEntry)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		writer, _ := zstd.NewWriter(&buf)
		_, _ = writer.Write(data)
		writer.Close()
	}
}

// Run Size Comparisons
func main() {
	logEntry := newLogEntry()
	protoLogEntry := newProtoLogEntry()

	// Serialize logs
	jsonData, _ := json.Marshal(logEntry)
	protoData, _ := proto.Marshal(protoLogEntry)
	msgpackData, _ := msgpack.Marshal(logEntry)

	// Compress logs using Zstd
	var jsonCompressed, msgpackCompressed, protoCompressed bytes.Buffer
	jsonWriter, _ := zstd.NewWriter(&jsonCompressed)
	msgpackWriter, _ := zstd.NewWriter(&msgpackCompressed)
	protoWriter, _ := zstd.NewWriter(&protoCompressed)

	jsonWriter.Write(jsonData)
	protoWriter.Write(protoData)
	msgpackWriter.Write(msgpackData)

	jsonWriter.Close()
	protoWriter.Close()
	msgpackWriter.Close()

	// Print file size results
	fmt.Println("Size Comparison (Bytes):")
	fmt.Printf("JSON: %d bytes\n", len(jsonData))
	fmt.Printf("Protobuf: %d bytes\n", len(protoData))
	fmt.Printf("MessagePack: %d bytes\n", len(msgpackData))
	fmt.Printf("JSON + Zstd: %d bytes\n", jsonCompressed.Len())
	fmt.Printf("Protobuf + Zstd: %d bytes\n", protoCompressed.Len())
	fmt.Printf("MessagePack + Zstd: %d bytes\n", msgpackCompressed.Len())
}
