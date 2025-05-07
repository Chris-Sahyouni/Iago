# Iago

Iago is a tool for crafting Return-Oriented-Programming payloads.

## Features
- Generates a complete ROP chain in a single command
- Finds individual gadgets in the target binary
- Adds padding bytes to ROP chains for buffer overflows
- **Supported file formats**: `ELF`
- **Supported ISAs**: `x86`, `x64`, `ARM`, `ARM (Thumb Mode)`, `AArch64`



## Limitations
- For ISAs with variable length instruction encodings such as `x86`, Iago may include gadgets that begin execution from the middle of an instruction encoding resulting in a different set of instructions being executed than intended.
- Iago only treats `ARM` binaries as entirely Thumb mode or entirely ARM. Because of this it can make mistakes parsing `ARM` binaries due to the fact that the size of instruction encodings can switch mid-execution between being either 4 bytes during ARM mode or either 2 or 4 bytes during Thumb mode.