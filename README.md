# Iago

Iago is a tool for crafting Return-Oriented-Programming payloads.

## Features
- Generates a complete ROP chain in a single command
- Finds individual gadgets in the target binary
- Adds padding bytes to ROP chains for buffer overflows
- **Supported file formats**: `ELF`
- **Supported ISAs**: `x86`, `x64`, `ARM`, `ARM (Thumb Mode)`, `AArch64`

## Installation

### Binary Releases
Binaries can be downloaded from the [releases](https://github.com/Chris-Sahyouni/iago/releases) page

### Using `go install`
Installing using `go install` requires Go 1.24.1 or later. To do so run
```
go install github.com/Chris-Sahyouni/iago@latest
```
### Compiling From Source
Compiling from source also requires Go 1.24.1 or later. First, install the necessary dependencies by running
```
go get
```
Then build the binary simply running
```
make
```

Note that compiling via the Makefile will not add the binary to `$GOPATH/bin`. If you want the binary to be available system-wide then `go install` would be the better installation method.

## Usage
Iago is an interactive shell, so to open it, after installing, simply run
```
iago
```
Then to specify the target binary run
```
load <path>
```
(Note: `iago` currently only targets ELF files)

### Finding Gadgets
Finding gadgets is as simple as running
```
find <gadget>
```
where `<gadget>` is a hexadecimal string representing the machine code of the target gadget. If successful, `find` returns the virtual address of `<gadget>`.

Note that `find` will only search the target binary for a single contiguous gadget. If you are trying to find multiple gadgets whose result when chained together is `<gadget>` then `rop` would be the appropriate command.

### Generating ROP Chains

#### Specifying the Payload
To generate a ROP chain, first you must specify the target payload (i.e. the instructions you want to actually execute as a result of your ROP chain). To do so, run
```
set-target
```
to manually input the target payload. Or
```
set-target <path>
```
To specify a file containing the desired payload.
Similar to `find`, `set-target` expects a hexadecimal string representing the machine code of the desired payload.

*Tip:*
To get the machine code for the payload you want to execute, [godbolt.org](https://godbolt.org/) is a great resource.

#### Generating the Chain
Once you have specified a target binary as well as a target payload, generating a ROP chain can be in a single command with
```
rop
```
The chain of addresses will be written to a file called rop_chain or one you specify if you use the -o flag. It is written in raw binary so that the chain can be easily piped into the process you want to hijack.

#### Padding Payloads
To add padding to the current payload for buffer overflows run
```
pad <bytes>
```
which will generate a new paylaod with `<bytes>` number of bytes of padding prepended to it.


If you want to use a payload in `set-payload` that has been already been padded, the name of the file *must* include the substring "pad" directly followed by the number of bytes of padding (e.g. rop_chain_pad32) otherwise `iago` will misinterpret the padding as a part of the chain.



## Limitations
- For ISAs with variable length instruction encodings such as `x86`, Iago may include gadgets that begin execution from the middle of an instruction encoding resulting in a different set of instructions being executed than intended.
- Iago only treats `ARM` binaries as entirely Thumb mode or entirely ARM. Because of this it can make mistakes parsing `ARM` binaries due to the fact that the size of instruction encodings can switch mid-execution between being either 4 bytes during ARM mode or either 2 or 4 bytes during Thumb mode.
- Iago searches for direct string matches when searching for gadgets in `find` and generating ROP chains with `rop`. It is semantically unaware. Because of this, it's a good idea to try several semantically-equivalent gadgets/targets when using these commands.