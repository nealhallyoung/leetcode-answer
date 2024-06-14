//1
impl Solution {
    pub fn final_value_after_operations(operations: Vec<String>) -> i32 {
        let mut res=0;
        for s in operations{
            match s.as_str() {
                "X++"|"++X"=>res+=1,
                "X--"|"--X"=>res-=1,
                _=>()
            }
        }
        res
    }
}

//2
impl Solution {
    pub fn final_value_after_operations(operations: Vec<String>) -> i32 {
        operations.iter().fold(0, 
            |res,curr| (
                if curr.as_bytes()[1]==b'+'{
                    res+1
                }else{
                    res-1
                }
            )
        )
    }
}
