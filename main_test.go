package main

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/klauspost/compress/zstd"
	"google.golang.org/protobuf/proto"
)

// Benchmark JSON Serialization
func BenchmarkJSONSerialization(b *testing.B) {
	logEntry := newLogEntry()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(logEntry)
	}
}

// Benchmark JSON Deserialization
func BenchmarkJSONDeserialization(b *testing.B) {
	logEntry := newLogEntry()
	data, _ := json.Marshal(logEntry)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var log AuditLog
		_ = json.Unmarshal(data, &log)
	}
}

// Benchmark JSON + Zstd Compression
func BenchmarkJSONZstdCompression(b *testing.B) {
	logEntry := newLogEntry()
	data, _ := json.Marshal(logEntry)
	writer, _ := zstd.NewWriter(io.Discard)
	var dst []byte

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		dst = writer.EncodeAll(data, dst[:0])
		writer.Close()
	}
}

// Benchmark Protobuf Serialization
func BenchmarkProtobufSerialization(b *testing.B) {
	logEntry := newProtoLogEntry()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(logEntry)
	}
}

// Benchmark Protobuf Deserialization
func BenchmarkProtobufDeserialization(b *testing.B) {
	logEntry := newProtoLogEntry()
	data, _ := proto.Marshal(logEntry)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var log ProtoAuditLog
		_ = proto.Unmarshal(data, &log)
	}
}

// Benchmark Protobuf + Zstd Compression
func BenchmarkProtobufZstdCompression(b *testing.B) {
	logEntry := newProtoLogEntry()
	data, _ := proto.Marshal(logEntry)
	writer, _ := zstd.NewWriter(io.Discard)
	var dst []byte

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		dst = writer.EncodeAll(data, dst[:0])
		writer.Close()
	}
}

// Benchmark MessagePack Serialization
func BenchmarkMessagePackSerialization(b *testing.B) {
	logEntry := newLogEntry()
	var dst []byte
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		dst, _ = logEntry.MarshalMsg(dst[:0])
	}
}

// Benchmark MessagePack de-Serialization
func BenchmarkMessagePackDeserialization(b *testing.B) {
	logEntry := newLogEntry()
	dst, _ := logEntry.MarshalMsg(nil)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		logEntry.UnmarshalMsg(dst)
	}
}

func BenchmarkMessagePackTupleSerialization(b *testing.B) {
	e := *newLogEntry()
	logEntry := AuditLogTuple(e)
	var dst []byte
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		dst, _ = logEntry.MarshalMsg(dst[:0])
	}
}

// Benchmark MessagePack + Zstd Compression
func BenchmarkMessagePackZstdCompression(b *testing.B) {
	logEntry := newLogEntry()
	data, _ := logEntry.MarshalMsg(nil)
	writer, _ := zstd.NewWriter(io.Discard)
	var dst []byte

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		dst = writer.EncodeAll(data, dst[:0])
		writer.Close()
	}
}

// Benchmark MessagePack + Zstd Compression
func BenchmarkMessagePackTupleZstdCompression(b *testing.B) {
	logEntry := AuditLogTuple(*newLogEntry())
	data, _ := logEntry.MarshalMsg(nil)
	writer, _ := zstd.NewWriter(io.Discard)
	var dst []byte

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		dst = writer.EncodeAll(data, dst[:0])
		writer.Close()
	}
}
