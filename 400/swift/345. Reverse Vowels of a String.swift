class Solution {
    func reverseVowels(_ s: String) -> String {
        var array = Array(s)
        let vowels: [Character] = ["a", "e", "i", "o", "u", "A", "E", "I", "O", "U"]
        var i = 0
        var j = array.count - 1

        while i < j {
            while i < j && !vowels.contains(array[i]) {
                i += 1
            }

            while i < j && !vowels.contains(array[j]) {
                j -= 1
            }

            if i < j {
                let temp = array[i]
                array[i] = array[j]
                array[j] = temp
                i += 1
                j -= 1
            }
        }

        return String(array)
    }
}