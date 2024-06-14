class Solution{
    func scoreOfString(_ s: String) -> Int{
        var res=0
        for i in 1..<s.count{
            let char1=s[s.index(s.startIndex, offsetBy: i-1)].asciiValue!
            let char2=s[s.index(s.startIndex, offsetBy: i)].asciiValue!
            res+=abs(Int(char1)-Int(char2))
        }
        return res
    }
}
