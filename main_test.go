package main

import (
	"encoding/json"
	"flag"
	"io"
	"testing"

	"github.com/klauspost/compress/zstd"
	"google.golang.org/protobuf/proto"
)

var batchCount int

func init() {
	flag.IntVar(&batchCount, "batch", 1, "batch count")
}

type lister[T any] struct {
	items      []T
	batchCount int
	c          int
}

func (l *lister[T]) next() T {
	defer func() {
		l.c++
	}()
	if l.c == l.batchCount {
		l.c = 0
	}
	return l.items[l.c]
}

func (l *lister[T]) isBatchFull() bool {
	return l.c == l.batchCount
}

func newLister[T any](items []T, batchCount int) *lister[T] {
	return &lister[T]{
		items:      items,
		batchCount: batchCount,
	}
}

// Benchmark JSON Serialization
func BenchmarkJSONSerialization(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	flag.Parse()

	lister := newLister[AuditLog](logEntries, batchCount)
	var entries []AuditLog
	for i := 0; i < b.N; i++ {
		entries = append(entries, lister.next())
		if lister.isBatchFull() {
			_, _ = json.Marshal(AuditLogList{Entries: entries})
			entries = []AuditLog{}
		}
	}
	if len(entries) != 0 {
		_, _ = json.Marshal(AuditLogList{Entries: entries})
	}
}

// Benchmark JSON Deserialization
func BenchmarkJSONDeserialization(b *testing.B) {
	data, _ := json.Marshal(AuditLogList{
		Entries: logEntries,
	})
	b.ResetTimer()
	b.ReportAllocs()

	var logList AuditLogList
	_ = json.Unmarshal(data, &logList)
}

// Benchmark JSON + Zstd Compression
func BenchmarkJSONZstdCompression(b *testing.B) {
	flag.Parse()
	b.ResetTimer()
	b.StopTimer()
	lister := newLister[AuditLog](logEntries, batchCount)
	var entries []AuditLog
	writer, _ := zstd.NewWriter(io.Discard)

	for i := 0; i < b.N; i++ {
		entries = append(entries, lister.next())
		if lister.isBatchFull() {
			data, _ := json.Marshal(AuditLogList{Entries: entries})
			var dst []byte
			b.StartTimer()
			_ = writer.EncodeAll(data, dst[:0])
			writer.Close()
			b.StopTimer()
			b.ReportAllocs()
			entries = []AuditLog{}
		}
	}
	if len(entries) != 0 {
		data, _ := json.Marshal(AuditLogList{Entries: entries})
		var dst []byte
		b.StartTimer()
		_ = writer.EncodeAll(data, dst[:0])
		writer.Close()
		b.StopTimer()
	}
	b.ReportAllocs()
}

// Benchmark Protobuf Serialization
func BenchmarkProtobufSerialization(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	flag.Parse()

	lister := newLister[*ProtoAuditLog](protoLogEntries, batchCount)
	var entries []*ProtoAuditLog
	for i := 0; i < b.N; i++ {
		entries = append(entries, lister.next())
		if lister.isBatchFull() {
			_, _ = proto.Marshal(&ProtoAuditLogList{Entries: entries})
			entries = []*ProtoAuditLog{}
		}
	}
	if len(entries) != 0 {
		_, _ = proto.Marshal(&ProtoAuditLogList{Entries: entries})
	}
}

// Benchmark Protobuf Deserialization
func BenchmarkProtobufDeserialization(b *testing.B) {
	data, _ := proto.Marshal(&ProtoAuditLogList{
		Entries: protoLogEntries,
	})
	b.ResetTimer()
	b.ReportAllocs()

	var logList ProtoAuditLogList
	_ = proto.Unmarshal(data, &logList)
}

// Benchmark Protobuf + Zstd Compression
func BenchmarkProtobufZstdCompression(b *testing.B) {
	flag.Parse()
	b.ResetTimer()
	b.StopTimer()

	lister := newLister[*ProtoAuditLog](protoLogEntries, batchCount)
	var entries []*ProtoAuditLog
	writer, _ := zstd.NewWriter(io.Discard)

	for i := 0; i < b.N; i++ {
		entries = append(entries, lister.next())
		if lister.isBatchFull() {
			data, _ := proto.Marshal(&ProtoAuditLogList{Entries: entries})
			var dst []byte
			b.StartTimer()
			_ = writer.EncodeAll(data, dst[:0])
			writer.Close()
			b.StopTimer()
			b.ReportAllocs()
			entries = []*ProtoAuditLog{}
		}
	}
	if len(entries) != 0 {
		data, _ := proto.Marshal(&ProtoAuditLogList{Entries: entries})
		var dst []byte
		b.StartTimer()
		_ = writer.EncodeAll(data, dst[:0])
		writer.Close()
		b.StopTimer()
	}
	b.ReportAllocs()

}

func BenchmarkMessagePackSerialization(b *testing.B) {
	var dst []byte
	b.ResetTimer()
	b.ReportAllocs()
	flag.Parse()

	lister := newLister[AuditLog](logEntries, batchCount)
	var entries []AuditLog
	for i := 0; i < b.N; i++ {
		entries = append(entries, lister.next())
		if lister.isBatchFull() {
			list := AuditLogList{Entries: entries}
			dst, _ = list.MarshalMsg(dst[:0])
			entries = []AuditLog{}
		}
	}
	if len(entries) != 0 {
		list := AuditLogList{Entries: entries}
		_, _ = list.MarshalMsg(dst[:0])
	}
}

// Benchmark JSON Deserialization
func BenchmarkMessagePackDeserialization(b *testing.B) {
	list := AuditLogList{Entries: logEntries}
	dst, _ := list.MarshalMsg(nil)
	b.ResetTimer()
	b.ReportAllocs()
	list.UnmarshalMsg(dst)
}

func BenchmarkMessagePackZstdCompression(b *testing.B) {
	flag.Parse()
	b.ResetTimer()
	b.StopTimer()
	lister := newLister[AuditLog](logEntries, batchCount)
	var entries []AuditLog
	writer, _ := zstd.NewWriter(io.Discard)

	for i := 0; i < b.N; i++ {
		entries = append(entries, lister.next())
		if lister.isBatchFull() {
			list := AuditLogList{Entries: entries}
			data, _ := list.MarshalMsg(nil)
			var dst []byte
			b.StartTimer()
			_ = writer.EncodeAll(data, dst[:0])
			writer.Close()
			b.StopTimer()
			b.ReportAllocs()
			entries = []AuditLog{}
		}
	}
	if len(entries) != 0 {
		list := AuditLogList{Entries: entries}
		data, _ := list.MarshalMsg(nil)
		var dst []byte
		b.StartTimer()
		_ = writer.EncodeAll(data, dst[:0])
		writer.Close()
		b.StopTimer()
	}
	b.ReportAllocs()
}

func BenchmarkMessagePackTupleSerialization(b *testing.B) {
	var dst []byte
	b.ResetTimer()
	b.ReportAllocs()
	flag.Parse()

	lister := newLister[AuditLog](logEntries, batchCount)
	var entries []AuditLogTuple
	for i := 0; i < b.N; i++ {
		entries = append(entries, AuditLogTuple(lister.next()))
		if lister.isBatchFull() {
			auditLogTupleList := AuditLogTupleList{
				Entries: entries,
			}
			dst, _ = auditLogTupleList.MarshalMsg(dst[:0])
			entries = []AuditLogTuple{}
		}
	}
	if len(entries) != 0 {
		list := AuditLogTupleList{Entries: entries}
		_, _ = list.MarshalMsg(dst[:0])
	}
}

// Benchmark MessagePack + Zstd Compression
func BenchmarkMessagePackTupleZstdCompression(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	b.ReportAllocs()
	flag.Parse()

	lister := newLister[AuditLog](logEntries, batchCount)
	var entries []AuditLogTuple
	writer, _ := zstd.NewWriter(io.Discard)
	for i := 0; i < b.N; i++ {
		entries = append(entries, AuditLogTuple(lister.next()))
		if lister.isBatchFull() {
			auditLogTupleList := AuditLogTupleList{
				Entries: entries,
			}
			data, _ := auditLogTupleList.MarshalMsg(nil)
			var dst []byte
			b.StartTimer()
			_ = writer.EncodeAll(data, dst[:0])
			writer.Close()
			b.StopTimer()
			entries = []AuditLogTuple{}
		}
	}
	if len(entries) != 0 {
		list := AuditLogTupleList{Entries: entries}
		data, _ := list.MarshalMsg(nil)
		var dst []byte
		b.StartTimer()
		_ = writer.EncodeAll(data, dst[:0])
		writer.Close()
		b.StopTimer()
	}
}

// BenchmarkMessagePackTupleDeserialization Benchmark Msgpack tuple deserialization
func BenchmarkMessagePackTupleDeserialization(b *testing.B) {
	list := AuditLogTupleList{
		Entries: func() (entries []AuditLogTuple) {
			for _, logEntry := range logEntries {
				entries = append(entries, AuditLogTuple(logEntry))
			}
			return
		}(),
	}
	dst, _ := list.MarshalMsg(nil)
	b.ResetTimer()
	b.ReportAllocs()
	list.UnmarshalMsg(dst)
}
