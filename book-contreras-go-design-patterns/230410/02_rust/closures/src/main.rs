fn main() {
    println!("Hello, world!");

    fn add_n(n: i32) -> impl Fn(i32) -> i32 {

        // This way doesn't work
        // fn inner(m: i32) -> i32{
        //     m + n
        // }
        // inner
        move |m| { m + n }
    }

    let add_5 = add_n(5);

    let out1 = add_5(10);

    println!("out1: {}", out1);

}
