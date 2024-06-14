import Foundation

class Solution {
    func mergeAlternately(_ word1: String, _ word2: String) -> String {
        var result = ""
        var lastIndex = 0
        if word1.count <= word2.count {
            for (k,v) in word1.enumerated() {
                lastIndex = k
                let index = word2.index(word2.startIndex,offsetBy: k) 
                result = result + String(v) + String(word2[index])
            }
            result = result + String(word2[neoIndex(lastIndex+1,str: word2)...])
        }else {
            for (k,v) in word1.enumerated() {
                lastIndex = k
                let index = word1.index(word1.startIndex,offsetBy: k) 
                result = result + String(v) + String(word1[index])
            }
            result = result + String(word2[neoIndex(lastIndex+1,str: word1)...])
        }
        return result
    }
}

// func neoIndex(index1: Int,index2: Int,str:String) -> ClosedRange<String.Index>{
//     let one = str.index(str.startIndex,offsetBy: index1)
//     let two = str.index(str.startIndex,offsetBy: index2)
//     let range = one...two
//     return range
// }
func neoIndex(_ index: Int, str:String) -> String.Index{
    let one = str.index(str.startIndex,offsetBy: index)
    return one
}

var word1 = "abc"
var word2 = "pqyr"

let ob = Solution()

var str = ob.mergeAlternately(word1, word2)
print(str)

