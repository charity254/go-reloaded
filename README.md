# Text Modifier Tool

A Go program that processes text files and applies various transformations including number base conversions, case modifications, and punctuation formatting.

## Current Progress

- [x] Basic file I/O (read input file, write to output file)
- [x] Tokenizer (splits text into words, punctuation, and special markers)
- [x] Hexadecimal to decimal conversion
- [x] Binary to decimal conversion
- [x] Case transformations (uppercase, lowercase, capitalize) - single word
- [ ] Case transformations with count (up, 2), (low, 3), (cap, 6)
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
  
- **Case modifications**: Transforms word casing
  - `word (up)` → `WORD`
  - `WORD (low)` → `word`
  - `word (cap)` → `Word`

## Features To Be Implemented

- Case modifications with multiple word count
- Punctuation spacing corrections
- Quote handling
- Article correction (a → an before vowels/h)
```
