impl Solution {
    pub fn convert_temperature(celsius: f64) -> Vec<f64> {
        let K=celsius+273.15;
        let F=celsius*1.80+32.00;
        let res=vec![K,F];
        res
    }
}
