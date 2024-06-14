impl Solution {
    pub fn get_concatenation(nums: Vec<i32>) -> Vec<i32> {
        let mut arr=vec![];
        for i in 0..=1{
            for j in nums.iter(){
                arr.push(*j);
            }
        }
        arr
    }
}
