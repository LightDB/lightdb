package database

import (
	"encoding/json"
)

// Get retrieves the value of a key. Returns an error if the key is undefined.
func (db *DB) Get(key string) (interface{}, error) {
	h := hashOf(key)

	rec, err := db.index.Find(int(h))
	if err != nil {
		return nil, err
	}

	var data interface{}
	if err := json.Unmarshal(rec.Value, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// Set adds a new key value pair.
func (db *DB) Set(key string, value interface{}) error {
	h := hashOf(key)

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err := db.index.Insert(int(h), data); err != nil {
		return err
	}

	go db.WriteDisk(h, value)

	return nil
}

// Update updates a key's value. Returns an error if the key to be updated is
// undefined.
func (db *DB) Update(key string, value interface{}) error {
	h := hashOf(key)

	err := db.index.Delete(int(h))
	if err != nil {
		return err
	}

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err := db.index.Insert(int(h), data); err != nil {
		return err
	}

	return nil
}

// Delete deletes a key value pair from the specified key. Returns an error if
// the key is undefined.
func (db *DB) Delete(key string) error {
	h := hashOf(key)

	err := db.index.Delete(int(h))
	if err != nil {
		return err
	}

	return nil
}
