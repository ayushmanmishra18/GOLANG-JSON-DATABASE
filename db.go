package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type DB struct {
	dir string
	mu  sync.RWMutex
}

func New(dir string, options interface{}) (*DB, error) {
	// create directory if not exists
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return &DB{dir: dir}, nil
}

// Write data
func (db *DB) Write(collection, key string, value interface{}) error {
	db.mu.Lock()         // 🔒 lock
	defer db.mu.Unlock() // 🔓 unlock

	filePath := filepath.Join(db.dir, collection+"_"+key+".json")

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

// Read all records
func (db *DB) Readall(collection string) ([]string, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	files, err := filepath.Glob(filepath.Join(db.dir, collection+"_*"))
	if err != nil {
		return nil, err
	}

	var records []string

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Read error:", err)
			continue
		}
		records = append(records, string(data))
	}

	return records, nil
}

// Read single record
func (db *DB) Read(collection, key string) (string, error) {
	db.mu.RLock()         // read lock
	defer db.mu.RUnlock()

	filePath := filepath.Join(db.dir, collection+"_"+key+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Update record (same as write, overwrite)
func (db *DB) Update(collection, key string, value interface{}) error {
	return db.Write(collection, key, value)
}

// Delete record
func (db *DB) Delete(collection, key string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	filePath := filepath.Join(db.dir, collection+"_"+key+".json")
	return os.Remove(filePath)
}