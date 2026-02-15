# Text Modifier Tool

A Go program that processes text files and applies various transformations including number base conversions, case modifications, and punctuation formatting.

## Current Progress

- [x] Basic file I/O (read input file, write to output file)
- [x] Tokenizer (splits text into words, punctuation, and special markers)
- [x] Hexadecimal to decimal conversion (single and multiple words)
- [x] Binary to decimal conversion (single and multiple words)
- [x] Case transformations (uppercase, lowercase, capitalize)
- [x] Case transformations with count (up, 2), (low, 3), (cap, 6)
- [ ] Punctuation spacing rules
- [ ] Quote handling
- [ ] Article correction (a/an)

## Usage
```bash
go run . input.txt output.txt
```

## Features Implemented

- **Number base conversions**: Converts hexadecimal and binary numbers to decimal
  - `1E (hex)` → `30`
  - `1010 (bin)` → `10`
  - `FF (hex, 2) AB` → `255 171`
  
- **Case modifications**: Transforms word casing
  - `word (up)` → `WORD`
  - `WORD (low)` → `word`
  - `word (cap)` → `Word`
  - `this is exciting (up, 2)` → `this IS EXCITING`
  - `IT WAS THE (low, 3) end` → `it was the end`

## Features To Be Implemented

- Punctuation spacing corrections
- Quote handling
- Article correction (a → an before vowels/h)

