use clap::Parser;

#[derive(Debug, Parser)]
pub struct Args {
    #[arg(short, long)]
    pub file: String,
}

pub struct Elf {

}

impl Elf {
    fn from(path: String) {
        
    }
}