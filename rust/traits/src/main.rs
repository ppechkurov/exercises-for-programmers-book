pub trait Summary {
    fn summarize(&self) -> String {
        format!("this is summary")
    }
}

struct Article {
    pub author: String,
}

impl Summary for Article {
    fn summarize(&self) -> String {
        format!("this is summary impl. Author: {}.", self.author)
    }
}

fn print_summary<T: Summary>(item: &T) {
    println!("trait summary: {}", item.summarize())
}

fn main() {
    let a = Article {
        author: String::from("E.M. Remark"),
    };

    println!("{}", a.summarize());
    print_summary(&a)
}
