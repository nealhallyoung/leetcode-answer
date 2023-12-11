class Solution {
    func kidsWithCandies(_ candies: [Int], _ extraCandies: Int) -> [Bool] {
        var maxCandy = 0
        var reslut: [Bool] = []
        for item in candies {
            if item > maxCandy {
                maxCandy = item
            }
        }
        for item in candies {
            if item + extraCandies >= maxCandy{
                reslut.append(true)
            }else {
                reslut.append(false)
            }
        }
        return reslut
    }
}