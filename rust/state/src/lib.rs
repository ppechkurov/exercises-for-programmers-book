pub struct Post {
    content: String,
}

impl Post {
    pub fn new() -> DraftPost {
        DraftPost {
            content: String::new(),
        }
    }

    pub fn content(&self) -> &str {
        &self.content
    }
}

pub struct DraftPost {
    content: String,
}

impl DraftPost {
    pub fn add_text(mut self, text: &str) -> Self {
        self.content.push_str(text);
        self
    }

    pub fn request_review(self) -> PendingReviewPost {
        PendingReviewPost {
            count: 0,
            content: self.content,
        }
    }
}

pub struct PendingReviewPost {
    count: u8,
    content: String,
}

impl PendingReviewPost {
    pub fn approve(self) -> Post {
        Post {
            content: self.content,
        }
    }
}
