# Iago

Iago is a tool for crafting Return-Oriented-Programming payloads.

## Features
- Generates a complete ROP chain in a single command
- Finds individual gadgets in the target binary
- Adds padding bytes to ROP chains for buffer overflows
- **Supported file formats**: `ELF`
- **Supported ISAs**: `x86`, `x64`, `ARM`

## Limitations
- For ISAs with variable length instruction encodings such as `x86`, Iago may include gadgets that begin execution from the middle of an instruction encoding resulting in a different set of instructions being executed than intended.