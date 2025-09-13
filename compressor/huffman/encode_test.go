package huffman

import (
	"reflect"
	"testing"
)

// TestBuildPrefixTable tests the buildPrefixTable function with a sample frequency map.
func TestBuildPrefixTable(t *testing.T) {
	// Example frequency map for testing
	frequency := map[rune]int{
		'a': 5,
		'b': 9,
		'c': 12,
		'd': 13,
		'e': 16,
		'f': 45,
	}

	// Expected prefix table (codes may vary, as Huffman trees can have multiple valid structures)
	expected := map[rune]string{
		'a': "1100",
		'b': "1101",
		'c': "100",
		'd': "101",
		'e': "111",
		'f': "0",
	}

	prefixTable := buildPrefixTable(frequency)

	// Verify the prefix table generated
	if !reflect.DeepEqual(prefixTable, expected) {
		t.Errorf("Prefix table does not match expected output.\nGot: %v\nExpected: %v", prefixTable, expected)
	}
}

// TestBuildPrefixTableEmpty tests the function with an empty frequency map.
func TestBuildPrefixTableEmpty(t *testing.T) {
	frequency := map[rune]int{}

	// Call the function
	prefixTable := buildPrefixTable(frequency)

	// Verify the result is an empty map
	if len(prefixTable) != 0 {
		t.Errorf("Expected an empty prefix table, got: %v", prefixTable)
	}
}

// TestBuildPrefixTableSingleElement tests the function with a map containing only one character.
func TestBuildPrefixTableSingleElement(t *testing.T) {
	frequency := map[rune]int{
		'a': 1,
	}

	// Expected prefix table (only one character, so it should get an empty code or "0")
	expected := map[rune]string{
		'a': "",
	}

	prefixTable := buildPrefixTable(frequency)

	// Verify the prefix table generated
	if !reflect.DeepEqual(prefixTable, expected) {
		t.Errorf("Prefix table does not match expected output.\nGot: %v\nExpected: %v", prefixTable, expected)
	}
}
