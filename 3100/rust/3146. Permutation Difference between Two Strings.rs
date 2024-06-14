impl Solution {
    pub fn find_permutation_difference(s: String, t: String) -> i32 {
        let mut res=0;
        for (i,cs) in s.char_indices() {
            for (j,ct) in t.char_indices() {
                if cs==ct {
                    res+=(i as i32 - j as i32).abs()
                }
            }
        }
        res
    }
}
