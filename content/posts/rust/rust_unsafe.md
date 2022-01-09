```
fn main() {
    let mut d = String::from("aaaa");
    let d_len = d.len();
    { // 去掉 or 不去掉
        let mut e = String::wtih_capacity(d_len);
        unsafe {
            // 两个指针指向相同的堆内存d
            ptr::copy(&d, &mut e, 1);
        };
        println!("{:?}", e.as_ptr());
        mem::drop(e);
    }
    println!("{:?}", d.as_ptr());
    d.push_str("a");
    println!("{:?}", d)
}
```