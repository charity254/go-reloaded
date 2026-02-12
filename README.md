# Text Modifier Tool

A Go program that processes text files and applies various transformations including number base conversions, case modifications, and punctuation formatting.

## Current Progress

- [x] Basic file I/O (read input file, write to output file)
- [x] Tokenizer (splits text into words, punctuation, and special markers)
- [ ] Hexadecimal to decimal conversion
- [ ] Binary to decimal conversion
- [ ] Case transformations (uppercase, lowercase, capitalize)
- [ ] Punctuation spacing rules
- [ ] Article correction (a/an)

## Usage
```bash
go run . input.txt output.txt
```

## Features (To Be Implemented)

- Number base conversions (hex, bin to decimal)
- Case modifications (up, low, cap) with optional count
- Punctuation spacing corrections
- Quote handling
- Article correction (a → an before vowels/h)
```