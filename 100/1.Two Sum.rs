struct Solution{

}
fn main() {
    let a=vec![3,2,4];
    let result=Solution::two_sum(a, 6);
    println!("{:?}",result);
}
//1
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        for i in 0..nums.len() {
            for j in i+1..nums.len() {
                if nums[i]+nums[j]==target {
                    return vec![i as i32,j as i32];
                }
            }
        }
        unreachable!()
    }
}

// 2
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut hashtable=HashMap::new();
        for (i,&v) in nums.iter().enumerate() {
            if let Some(&j)=hashtable.get(&(target-v)){
                return vec![i as i32,j as i32]
            }
            hashtable.insert(v,i);
        }
        unreachable!()
    }
}
