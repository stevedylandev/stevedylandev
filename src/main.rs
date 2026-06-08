use std::process::Command;
use std::{thread, time::Duration};

struct Human {
    name: &'static str,
    title: &'static str,
    job: &'static str,
    website: &'static str,
    dream: &'static str,
    hobbies: &'static [&'static str],
    side_projects: &'static [&'static str],
}

impl Human {
    fn introduce(&self) {
        println!("Name: {}", self.name);
        println!("Title: {}", self.title);
        println!("Job: {}", self.job);
        println!("Website: {}", self.website);
        println!("Hobbies:");
        for hobby in self.hobbies {
            println!("  - {hobby}");
        }
        println!("Side Projects:");
        for project in self.side_projects {
            println!("  - {project}");
        }
        println!("Dream: {}", self.dream);
    }

    fn surprise(&self) -> String {
        let hex = "6375726c2061736369692e6c6976652f7269636b";
        let bytes: Vec<u8> = (0..hex.len())
            .step_by(2)
            .map(|i| u8::from_str_radix(&hex[i..i + 2], 16).unwrap())
            .collect();
        String::from_utf8(bytes).unwrap()
    }
}

fn main() {
    let steve = Human {
        name: "Steve Simkins",
        title: "Senior Solutions Engineer",
        job: "https://stablecore.com",
        website: "https://stevedylan.dev",
        dream: "Help build and promote an open web",
        hobbies: &["photography", "bird watching", "blogging", "building personal software"],
        side_projects: &["https://andromeda.build", "https://sequoia.pub", "https://bhvr.dev"],
    };

    steve.introduce();
    thread::sleep(Duration::from_secs(2));

    Command::new("sh")
        .arg("-c")
        .arg(steve.surprise())
        .status()
        .expect("Failed to run surprise :(");
}
