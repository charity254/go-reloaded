package main

import "testing"

func TestHex(t *testing.T) {
	//TEST 1
	input := []string {"1E", "(hex)", "files"}
	expected := []string {"30", "files"}
	result := hex(input)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i:= range expected {
		if result[i] != expected[i] {
			t.Errorf("At index %d: expected %s, got %s", i, expected[i], result[i])
		}
	}
//TEST 2
	input2 := []string{"FF", "AB", "(hex, 2)", "values"}
	expected2 := []string{"255", "171", "values"}
	result2 := hex(input2)

	if len(result2) != len(expected2) {
		t.Errorf("Test 2 - Expected length %d, got %d", len(expected2), len(result2))
	}
	for i := range expected2 {
		if result2[i] != expected2[i] {
			t.Errorf("Test 2 - At index %d: expected %s, got %s", i, expected2[i], result2[i])
		}
	}
}

func TestBin(t *testing.T) {
	//TEST 1
	input := []string{"10", "(bin)", "years"}
	expected := []string{"2", "years"}
	result := bin(input)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i:= range expected {
		if result[i] != expected[i] {
			t.Errorf("At index %d: expected %s, got %s", i, expected[i], result[i])
		}
	}
//TEST 2
	input2 := []string{"1010", "1111", "(bin, 2)", "numbers"}
	expected2 := []string{"10", "15", "numbers"}
	result2 := bin(input2)

	if len(result2) != len(expected2) {
		t.Errorf("Test 2 - Expected length %d, got %d", len(expected2), len(result2))
	}
	for i := range expected2 {
		if result2[i] != expected2[i] {
			t.Errorf("Test 2 - At index %d: expected %s, got %s", i, expected2[i], result2[i])
		}
	}
}


func TestUpper(t *testing.T) {
	//TEST 1
	input := []string{"ready", "set", "go", "(up)"}
	expected := []string{"ready", "set", "GO"}
	result := upper(input)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i:= range expected {
		if result[i] != expected[i] {
			t.Errorf("At index %d: expected %s, got %s", i, expected[i], result[i])
		}
	}
//TEST 2
	input2 := []string{"this", "is", "so", "exciting", "(up, 2)"}
	expected2 := []string{"this", "is", "SO", "EXCITING"}
	result2 := upper(input2)

	if len(result2) != len(expected2) {
		t.Errorf("Test 2 - Expected length %d, got %d", len(expected2), len(result2))
	}
	for i := range expected2 {
		if result2[i] != expected2[i] {
			t.Errorf("Test 2 - At index %d: expected %s, got %s", i, expected2[i], result2[i])
		}
	}
}

func TestLower(t *testing.T) {
	//TEST 1
	input := []string{"I", "should", "stop", "SHOUTING", "(low)"}
	expected := []string{"I", "should", "stop", "shouting"}
	result := lower(input)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i:= range expected {
		if result[i] != expected[i] {
			t.Errorf("At index %d: expected %s, got %s", i, expected[i], result[i])
		}
	}
//TEST 2
	input2 := []string{"IT", "WAS", "THE", "(low, 3)", "end"}
	expected2 := []string{"it", "was", "the", "end"}
	result2 := lower(input2)

	if len(result2) != len(expected2) {
		t.Errorf("Test 2 - Expected length %d, got %d", len(expected2), len(result2))
	}
	for i := range expected2 {
		if result2[i] != expected2[i] {
			t.Errorf("Test 2 - At index %d: expected %s, got %s", i, expected2[i], result2[i])
		}
	}
}

func TestCap(t *testing.T) {
	//TEST 1
	input := []string{"welcome", "to", "the", "bridge", "(cap)"}
	expected := []string{"welcome", "to", "the", "Bridge"}
	result := cap(input)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i:= range expected {
		if result[i] != expected[i] {
			t.Errorf("At index %d: expected %s, got %s", i, expected[i], result[i])
		}
	}
//TEST 2
	input2 := []string{"it", "was", "the", "age", "of", "foolishness", "(cap, 3)"}
	expected2 := []string{"it", "was", "the", "Age", "Of", "Foolishness"}
	result2 := cap(input2)

	if len(result2) != len(expected2) {
		t.Errorf("Test 2 - Expected length %d, got %d", len(expected2), len(result2))
	}
	for i := range expected2 {
		if result2[i] != expected2[i] {
			t.Errorf("Test 2 - At index %d: expected %s, got %s", i, expected2[i], result2[i])
		}
	}
}

func TestArticle(t *testing.T) {
	//TEST 1 (vowel lowercase)
	input := []string{"There", "was", "a", "apple"}
	expected := []string{"There", "was", "an", "apple"}
	result := article(input)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i:= range expected {
		if result[i] != expected[i] {
			t.Errorf("At index %d: expected %s, got %s", i, expected[i], result[i])
		}
	}
//TEST 2 (H- LOWERCASE)
	input2 := []string{"She", "is", "a", "honest", "person"}
	expected2 := []string{"She", "is", "an", "honest", "person"}
	result2 := article(input2)

	if len(result2) != len(expected2) {
		t.Errorf("Test 2 - Expected length %d, got %d", len(expected2), len(result2))
	}
	for i := range expected2 {
		if result2[i] != expected2[i] {
			t.Errorf("Test 2 - At index %d: expected %s, got %s", i, expected2[i], result2[i])
		}
	}
	//TEST 3 (vowel uppercase)
	input3 := []string{"A", "amazing", "day"}
	expected3 := []string{"An", "amazing", "day"}
	result3 := article(input3)

	if len(result3) != len(expected3) {
		t.Errorf("Test 3 - Expected length %d, got %d", len(expected3), len(result3))
	}
	for i := range expected3 {
		if result3[i] != expected3[i] {
			t.Errorf("Test 3 - At index %d: expected %s, got %s", i, expected3[i], result3[i])
		}
	}
	//TEST 4(No change)
	input4 := []string{"I", "saw", "a", "cat"}
	expected4 := []string{"I", "saw", "a", "cat"}
	result4 := article(input4)

	if len(result4) != len(expected4) {
		t.Errorf("Test 4 - Expected length %d, got %d", len(expected4), len(result4))
	}
	for i := range expected4 {
		if result4[i] != expected4[i] {
			t.Errorf("Test 4 - At index %d: expected %s, got %s", i, expected4[i], result4[i])
		}
	}
}

func TestPunctuation(t *testing.T) {
	//TEST 1 Basic punctuation spacing
	input := []string{"Hello", ",", "world", "!"}
	expected := "Hello, world!"
	result := punctuation(input)

	if result != expected {
    t.Errorf("Test 2 - Expected '%s', got '%s'", expected, result)
}
//TEST 2 Grouped punctuation
	input2 := []string{"Wait", ".", ".", ".", "what"}
	expected2 := "Wait... what"
	result2 := punctuation(input2)

	if result2 != expected2 {
    t.Errorf("Test 4 - Expected '%s', got '%s'", expected2, result2)
}
	//TEST 3 Multiple grouped punctuation
	input3 := []string{"Really", "!", "?"}
	expected3 := "Really!?"
	result3 := punctuation(input3)

	if result3!= expected3{
    t.Errorf("Test 3 - Expected '%s', got '%s'", expected3, result3)
}
	//TEST 4 Quotes
	input4 := []string{"I", "am", "'", "awesome", "'"}
	expected4 :=  "I am 'awesome'"
	result4 := punctuation(input4)

	if result4 != expected4 {
    t.Errorf("Test 4 - Expected '%s', got '%s'", expected4, result4)
 }
 	//TEST 5 Multi-word Quotes
	input5 := []string{"I", "am", "'", "awesome", "'"}
	expected5 :=  "I am 'awesome'"
	result5 := punctuation(input5)

	if result4 != expected4 {
    t.Errorf("Test 5 - Expected '%s', got '%s'", expected5, result5)
 }
}