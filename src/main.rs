use Jaws::Args;
use clap::Parser;
use std::{path::Path, fs::File, io::prelude::*};


fn main() {
    let args: Args = Args::parse();

    let file: File = File::open(args.file).unwrap();



}

// segment: information for runtime execution
// sections: information for linking and relocation

/* -------------------------------- ELF FILES ------------------------------- */

/*
ELF HEADER


*/