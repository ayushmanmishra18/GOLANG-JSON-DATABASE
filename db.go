package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type DB struct {
	dir string
	mu  sync.RWMutex
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address string
}

// Create DB
func New(dir string, options interface{}) (*DB, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return &DB{dir: dir}, nil
}

// Write
func (db *DB) Write(collection, key string, value interface{}) error {
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	key = strings.ReplaceAll(key, "/", "_")

	db.mu.Lock()
	defer db.mu.Unlock()

	filePath := filepath.Join(db.dir, collection+"_"+key+".json")

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

// Read all
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
			continue
		}
		records = append(records, string(data))
	}

	return records, nil
}

// Read one
func (db *DB) Read(collection, key string) (string, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	filePath := filepath.Join(db.dir, collection+"_"+key+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Update
func (db *DB) Update(collection, key string, value interface{}) error {
	return db.Write(collection, key, value)
}

// Delete
func (db *DB) Delete(collection, key string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	filePath := filepath.Join(db.dir, collection+"_"+key+".json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("record not found")
	}

	return os.Remove(filePath)
}