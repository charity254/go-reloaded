# Text Modifier Tool

A Go program that processes text files and applies various transformations including number base conversions, case modifications, and punctuation formatting.

## Current Progress

- [x] Basic file I/O (read input file, write to output file)
- [x] Tokenizer (splits text into words, punctuation, and special markers)
- [x] Hexadecimal to decimal conversion (single and multiple words)
- [x] Binary to decimal conversion (single and multiple words)
- [x] Case transformations (uppercase, lowercase, capitalize)
- [x] Case transformations with count (up, 2), (low, 3), (cap, 6)
- [x] Punctuation spacing rules
- [x] Quote handling
- [x] Article correction (a/an)

## Usage
```bash
go run . input.txt output.txt
```

## Features Implemented

### Number Base Conversions
Converts hexadecimal and binary numbers to decimal:
- `1E (hex)` → `30`
- `1010 (bin)` → `10`
- `FF AB (hex, 2)` → `255 171`

### Case Modifications
Transforms word casing:
- `word (up)` → `WORD`
- `WORD (low)` → `word`
- `word (cap)` → `Word`
- `this is exciting (up, 2)` → `this IS EXCITING`

### Punctuation Spacing
Fixes spacing around punctuation marks (`.`, `,`, `!`, `?`, `:`, `;`):
- `Hello , world !` → `Hello, world!`
- `Wait ... what` → `Wait... what`
- `Really !?` → `Really!?`

### Quote Handling
Properly formats single quotes around words:
- `I am ' awesome '` → `I am 'awesome'`
- `He said : ' I am the best '` → `He said: 'I am the best'`

### Article Correction
Converts `a` to `an` before vowels (a, e, i, o, u) or h:
- `a apple` → `an apple`
- `a honest person` → `an honest person`
- `A amazing day` → `An amazing day`

## Project Structure
```
.
├── main.go           # Entry point and file I/O
├── tokenize.go       # Text tokenization
├── processor.go      # All transformation functions
└── README.md
```
