class Solution {
    func maximumHappinessSum(_ happiness: [Int], _ k: Int) -> Int {
        var sortHappiness = happiness.sorted(by:>)
        var ans: Int64 = 0
        var i=0
        for _ in 0..<min(k,sortHappiness.count){
            let temp = max(sortHappiness[i]-i,0)
            ans += Int64(temp)
            i += 1
        }
        return Int(ans)
    }
}