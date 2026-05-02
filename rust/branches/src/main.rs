fn main() {
    let condition = false;
    let number = if condition {
        let n = 5;
        n
    } else {
        "six"
    };
    println!("The value of number is: {number}")
}
