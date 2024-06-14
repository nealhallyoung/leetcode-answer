class Solution {
    func scoreOfString(_ s: String) -> Int {
        var res = 0
        for i in 0..<s.count - 1 {
            let char1 = s[s.index(s.startIndex, offsetBy: i)].asciiValue!
            let char2 = s[s.index(s.startIndex, offsetBy: i + 1)].asciiValue!
            res += abs(Int(char1) - Int(char2))
        }
        return res
    }
}