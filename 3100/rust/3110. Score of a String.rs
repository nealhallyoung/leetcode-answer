impl Solution {
    pub fn score_of_string(s: String) -> i32 {
        let mut result=0;
        let mut value=vec![];
        let mut my_string=s.chars();
        while let Some(ch) = my_string.next() {
            println!("{}",ch);
            if ch.is_ascii() {
                value.push(ch as i32)
            }
        }
        for i in 1..value.len() {
            result+=(value[i-1]-value[i]).abs();
        }
        result
    }
}