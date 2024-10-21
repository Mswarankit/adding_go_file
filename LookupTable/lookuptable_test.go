package main

import (
	"testing"
)

func TestSet(t *testing.T) {
	table := New()
	table.Set("Name", "Dyna Pyssan")

	if value, exists := table.Get("Name"); !exists || value != "Dyna" {
		t.Errorf("Expected 'Dyna', got '%v'", value)
	}
}

func TestGet(t *testing.T) {
	table := New()
	table.Set("Age", 30)

	if value, exists := table.Get("Age"); !exists || value != 30 {
		t.Errorf("Expected 30, got %v", value)
	}

	// Testing a non-existing key
	if _, exists := table.Get("NonExisting"); exists {
		t.Error("Expected no value for 'NonExisting' key")
	}
}

func TestLen(t *testing.T) {
	table := New()

	if table.Len() != 0 {
		t.Error("Expected length 0")
	}

	table.Set("Key", "Value")
	if table.Len() != 1 {
		t.Error("Expected length 1")
	}

	table.Set("Key2", "Value2")
	if table.Len() != 2 {
		t.Error("Expected length 2")
	}

	table.Delete("Key")
	if table.Len() != 1 {
		t.Error("Expected length 1 after deletion")
	}
}

func TestDelete(t *testing.T) {
	table := New()
	table.Set("Name", "Dyna")
	table.Delete("Name")

	if _, exists := table.Get("Name"); exists {
		t.Error("Expected 'Name' key to be deleted")
	}

	// Trying to delete a non-existing key should not panic
	table.Delete("NonExisting")
}

func TestKeys(t *testing.T) {
	table := New()
	table.Set("Key1", "Value1")
	table.Set("Key2", "Value2")

	keys := table.Keys()
	if len(keys) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(keys))
	}

	keySet := make(map[any]bool)
	for _, key := range keys {
		keySet[key] = true
	}

	if !keySet["Key1"] || !keySet["Key2"] {
		t.Errorf("Expected keys to be 'Key1' and 'Key2', got %v", keys)
	}
}

func TestValues(t *testing.T) {
	table := New()
	table.Set("Key1", "Value1")
	table.Set("Key2", "Value2")

	values := table.Values()
	if len(values) != 2 {
		t.Errorf("Expected 2 values, got %d", len(values))
	}

	valueSet := make(map[any]bool)
	for _, value := range values {
		valueSet[value] = true
	}

	if !valueSet["Value1"] || !valueSet["Value2"] {
		t.Errorf("Expected values to be 'Value1' and 'Value2', got %v", values)
	}
}

func TestClear(t *testing.T) {
	table := New()
	table.Set("Key", "Value")
	table.Clear()

	if table.Len() != 0 {
		t.Error("Expected length 0 after clearing")
	}
}

func TestContains(t *testing.T) {
	table := New()
	table.Set("Key", "Value")

	if !table.Contains("Key") {
		t.Error("Expected 'Key' to be present")
	}

	if table.Contains("NonExisting") {
		t.Error("Expected 'NonExisting' to not be present")
	}
}

func TestUpdate(t *testing.T) {
	table1 := New()
	table1.Set("Key1", "Value1")
	table2 := New()
	table2.Set("Key2", "Value2")

	table1.Update(table2)

	if value, exists := table1.Get("Key2"); !exists || value != "Value2" {
		t.Errorf("Expected 'Value2', got %v", value)
	}
}

func TestDefault(t *testing.T) {
	table := New()
	table.Set("Key", "Value")

	// Testing existing key
	if result := table.Default("Key", "DefaultValue"); result != "Value" {
		t.Errorf("Expected 'Value', got %v", result)
	}

	// Testing non-existing key
	if result := table.Default("NonExisting", "DefaultValue"); result != "DefaultValue" {
		t.Errorf("Expected 'DefaultValue', got %v", result)
	}
}

func TestSetDefaultValue(t *testing.T) {
	table := New()

	if result := table.SetDefaultValue("Key", "DefaultValue"); result != "DefaultValue" {
		t.Errorf("Expected 'DefaultValue', got %v", result)
	}

	// Setting an existing key
	table.Set("Key", "NewValue")
	if result := table.SetDefaultValue("Key", "DefaultValue"); result != "NewValue" {
		t.Errorf("Expected 'NewValue', got %v", result)
	}
}

func TestPop(t *testing.T) {
	table := New()
	table.Set("Key", "Value")

	// Pop an existing key
	if value, err := table.Pop("Key"); err != nil || value != "Value" {
		t.Errorf("Expected 'Value', got %v, err: %v", value, err)
	}

	// Pop a non-existing key
	if _, err := table.Pop("NonExisting"); err == nil {
		t.Error("Expected error for non-existing key")
	}
}

func TestIterator(t *testing.T) {
	table := New()
	table.Set("Key1", "Value1")
	table.Set("Key2", "Value2")

	count := 0
	for kv := range table.Iterator() {
		count++
		if kv.Key == "Key1" && kv.Value != "Value1" {
			t.Errorf("Expected 'Value1' for 'Key1', got %v", kv.Value)
		}

		if kv.Key == "Key2" && kv.Value != "Value2" {
			t.Errorf("Expected 'Value2' for 'Key2', got %v", kv.Value)
		}
	}

	if count != 2 {
		t.Errorf("Expected 2 items iterated, got %d", count)
	}
}
