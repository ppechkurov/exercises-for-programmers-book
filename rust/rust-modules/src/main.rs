mod garden;

fn main() {
    let plant = garden::vegetables::Asparagus {
        name: String::from("Asparagus"),
        stalks: 5,
    };

    println!("I'm growing {:?}!", plant);
}
