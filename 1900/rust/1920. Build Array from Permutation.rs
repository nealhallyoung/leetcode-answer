impl Solution {
    pub fn build_array(nums: Vec<i32>) -> Vec<i32> {
        let mut res=vec![];
        for i in 0..nums.len() {
            res.insert(i, nums[nums[i] as usize])
        }
        res
    }
}
