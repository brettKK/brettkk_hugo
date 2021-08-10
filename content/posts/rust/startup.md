Rust程序设计

+ 安装

https://rustup.rs/

+ 第一个程式


cargo -h
cargo new --bin hello
cargo run 

rustup doc
cargo doc --open

std::collections
Vec, VecDeque, HashMap, String


## 编写测试

rust中的测试是带有标注test属性的函数。 #[test]

cargo test.

## 智能指针

Box<T>

Rc<T>

RefCell<T>

Deref, Drop trait

Weak<T>
