```bash
git clone https://github.com/stevedylandev/stevedylandev && cd stevedylandev && cargo run
```

```rust
use std::{thread, time::Duration};
use std::process::Command;

struct Human {
    name: String,
    title: String,
    job: String,
    hobbies: Vec<String>,
    dream: String,
    website: String,
    side_projects: Vec<String>,
}

impl Human {
    fn introduce(&self) {
        println!("Name: {}", self.name);
        println!("Title: {}", self.title);
        println!("Job: {}", self.job);
        println!("Website: {}", self.website);

        println!("Hobbies:");
        for hobby in &self.hobbies {
            println!("  - {}", hobby);
        }

        println!("Side Projects:");
        for project in &self.side_projects {
            println!("  - {}", project);
        }

        println!("Dream: {}", self.dream);
    }

    fn surprise(&self) -> String {
        let hex = "6375726c2061736369692e6c6976652f7269636b";
        String::from_utf8(
            hex.chars()
                .collect::<Vec<char>>()
                .chunks(2)
                .map(|chunk| {
                    let hex_str: String = chunk.iter().collect();
                    u8::from_str_radix(&hex_str, 16).unwrap_or(0)
                })
                .collect::<Vec<u8>>(),
        )
        .unwrap_or_default()
    }
}

fn main() {
    let steve = Human {
        name: String::from("Steve Simkins"),
        title: String::from("DX Engineer"),
        job: String::from("https://namehash.io"),
        hobbies: vec![
            String::from("photography"),
            String::from("specialty coffee"),
            String::from("programming"),
        ],
        dream: String::from(
            "Help build the future of the web through developers tools, cryptography, and trust based protocols",
        ),
        website: String::from("https://stevedylan.dev"),
        side_projects: vec![
            String::from("https://orbiter.host"),
            String::from("https://bhvr.dev"),
        ],
    };

    steve.introduce();
    thread::sleep(Duration::from_secs(2));
    let secret = steve.surprise();
    let mut child = Command::new("sh")
        .arg("-c")
        .arg(&secret)
        .spawn()
        .expect("Failed");
    child.wait().expect("Failed");
}
```

