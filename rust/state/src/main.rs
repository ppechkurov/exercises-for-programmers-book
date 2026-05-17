use state::Post;

fn main() {
    let mut post = Post::new();

    let mut post = post
        .add_text("I ate a salad for lunch today")
        .add_text("Another text")
        .request_review()
        .approve();
}
