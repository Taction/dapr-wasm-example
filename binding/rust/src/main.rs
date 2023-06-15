use std::env;
use std::io;
use std::io::Write;

fn main() {
    let args: Vec<String> = env::args().collect();
    let mut stdout = io::stdout();
    let stdin = io::stdin();
    let mut buf = String::new();
    stdin.read_line(&mut buf).unwrap();
    // write args from index 1 to stdout
    for i in 1..args.len() {
        stdout.write_all(args[i].as_bytes()).unwrap();
        stdout.write_all(b" ").unwrap();
    }
    // write stdin to stdout
    stdout.write_all(buf.as_bytes()).unwrap();
    stdout.write_all(b"\n").unwrap();
}