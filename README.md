# Text Modifier Tool

A Go program that processes text files and applies various transformations including number base conversions, case modifications, and punctuation formatting.

## Description

This tool reads a text file, applies a series of modifications based on special markers and formatting rules, and outputs the processed text to a new file. It handles number base conversions, case transformations, punctuation spacing, quote formatting, and article corrections.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Testing](#testing)
- [Project Structure](#project-structure)
- [Examples](#examples)
- [Progress](#progress)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-reloaded.git
cd go-reloaded
```

2. Ensure you have Go installed (version 1.16 or higher):
```bash
go version
```

## Usage

Run the program with two arguments: input file and output file.
```bash
go run . input.txt output.txt
```

**Arguments:**
- First argument: Path to the input file containing text to be modified
- Second argument: Path to the output file where results will be written

**Example:**
```bash
go run . sample.txt result.txt
```

## Features

### 1. Number Base Conversions

#### Hexadecimal to Decimal
Converts hexadecimal numbers to decimal when followed by `(hex)`.

**Examples:**
- `1E (hex) files were added` → `30 files were added`
- `The value is FF (hex)` → `The value is 255`

**Multiple words:**
- `FF AB (hex, 2) values` → `255 171 values`

#### Binary to Decimal
Converts binary numbers to decimal when followed by `(bin)`.

**Examples:**
- `It has been 10 (bin) years` → `It has been 2 years`
- `Binary 1010 (bin) equals ten` → `Binary 10 equals ten`

**Multiple words:**
- `1010 1111 (bin, 2) numbers` → `10 15 numbers`

### 2. Case Transformations

#### Uppercase
Converts the word before `(up)` to uppercase.

**Examples:**
- `Ready, set, go (up)!` → `Ready, set, GO!`
- `This is so exciting (up, 2)` → `This is SO EXCITING`

#### Lowercase
Converts the word before `(low)` to lowercase.

**Examples:**
- `I should stop SHOUTING (low)` → `I should stop shouting`
- `IT WAS THE (low, 3) winter` → `it was the winter`

#### Capitalize
Capitalizes the word before `(cap)` (first letter uppercase, rest lowercase).

**Examples:**
- `Welcome to the brooklyn bridge (cap)` → `Welcome to the Brooklyn Bridge`
- `it was the age of foolishness (cap, 6)` → `It Was The Age Of Foolishness`

**Note:** For `(low)`, `(up)`, and `(cap)`, if a number appears like `(up, 2)`, it applies the transformation to the specified number of previous words.

### 3. Punctuation Spacing

Automatically fixes spacing around punctuation marks: `.` `,` `!` `?` `:` `;`

**Rules:**
- Punctuation should be close to the previous word (no space before)
- Space should appear after punctuation (before next word)
- Groups of punctuation like `...` or `!?` are kept together

**Examples:**
- `I was sitting over there ,and then BAMM !!` → `I was sitting over there, and then BAMM!!`
- `I was thinking ... You were right` → `I was thinking... You were right`
- `What !? Really ?` → `What!? Really?`

### 4. Quote Handling

Single quotes `'` are properly formatted around words with no internal spacing.

**Examples:**
- `I am exactly how they describe me: ' awesome '` → `I am exactly how they describe me: 'awesome'`
- `As he said: ' I am the most well-known person in the world '` → `As he said: 'I am the most well-known person in the world'`

### 5. Article Correction

Automatically converts `a` to `an` when the next word begins with a vowel (a, e, i, o, u) or h.

**Examples:**
- `There it was. A amazing rock!` → `There it was. An amazing rock!`
- `She is a honest person` → `She is an honest person`
- `A apple a day` → `An apple a day`
- `I saw a cat` → `I saw a cat` (no change)

**Note:** Preserves capitalization (A → An, a → an)

## Testing

The project includes comprehensive unit tests for all transformation functions.

### Run All Tests
```bash
go test
```

### Run Tests with Verbose Output
```bash
go test -v
```

### Run Specific Test
```bash
go test -run TestHex
go test -run TestPunctuation
```

### Test Coverage
```bash
go test -cover
```

## Project Structure
```
go-reloaded/
├── main.go              # Entry point, file I/O, and processing pipeline
├── tokenize.go          # Text tokenization into words and special markers
├── processor.go         # Transformation functions (hex, bin, upper, lower, cap, article, punctuation)
├── processor_test.go    # Unit tests for all transformation functions
├── README.md            # Project documentation
├── sample.txt           # Sample input file (optional)
└── result.txt           # Sample output file (optional)
```

## Examples

### Example 1: Basic Transformations

**Input (sample.txt):**
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom.
```

**Command:**
```bash
go run . sample.txt result.txt
```

**Output (result.txt):**
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom.
```

### Example 2: Number Conversions

**Input:**
```
I have 1E (hex) apples and 1010 (bin) oranges.
```

**Output:**
```
I have 30 apples and 10 oranges.
```

### Example 3: Combined Features

**Input:**
```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```

**Output:**
```
Simply add 66 and 2 and you will see the result is 68.
```

### Example 4: Complex Text

**Input:**
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```

**Output:**
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

## Progress

- [x] Basic file I/O (read input file, write to output file)
- [x] Tokenizer (splits text into words, punctuation, and special markers)
- [x] Hexadecimal to decimal conversion (single and multiple words)
- [x] Binary to decimal conversion (single and multiple words)
- [x] Case transformations (uppercase, lowercase, capitalize)
- [x] Case transformations with count (up, n), (low, n), (cap, n)
- [x] Punctuation spacing rules
- [x] Quote handling
- [x] Article correction (a/an)
- [x] Unit tests for all transformation functions

## Requirements

- Go 1.16 or higher
- Standard Go packages only (no external dependencies)

## Author

Your Name

## License

This project is part of a coding challenge/learning exercise.