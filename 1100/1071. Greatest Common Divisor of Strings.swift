class Solution {
    func gcdOfStrings(_ str1: String, _ str2: String) -> String {
        if str1 + str2 != str2 + str1 {
            return ""
        }
        let index = str1.index(str1.startIndex, offsetBy: gcd(a: str1.count,b: str2.count)-1, limitedBy: str1.endIndex)
        return String(str1[str1.startIndex...index!])
    }
}

func gcd(a:Int,b:Int) -> Int {
    var x = a
    var y = b
    if x<y {
        let temp = x
        x = y
        y = temp
    }
    var temp = x
    while temp != 0 {
        temp = x%y
        x = y
        y = temp
    }
    return x
}