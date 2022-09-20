#[macro_use] extern crate nickel;

use nickel::{Nickel, HttpRouter};

fn main() {
    let mut server = Nickel::new();

    server.get("/", middleware!("Hello World!"));
    server.get("/hello/:name", middleware! { |request|
        format!("Hello {}", request.param("name").unwrap())
    });

    server.listen("127.0.0.1:5000").unwrap();
}
