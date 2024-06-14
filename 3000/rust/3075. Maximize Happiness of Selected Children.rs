impl Solution {
    pub fn maximum_happiness_sum(happiness: Vec<i32>, k: i32) -> i64 {
        let mut happiness = happiness;
        happiness.sort_by(|a,b| b.cmp(a));
        let mut sum:i64=0;
        for i in 0..k{
            let temp:i32=happiness[i as usize] -i;
            sum+=if temp<=0 {0 as i64} else {temp as i64};
        }
        sum
    }
}