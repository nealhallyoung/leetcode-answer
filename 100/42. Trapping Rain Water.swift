class Solution {
  func trap(_ height: [Int]) -> Int {
    var (l, r) = (0, height.count - 1)
    var (lm, rm) = (height[l], height[r])
    var res = 0
    while l<r {
        if lm < rm{
                l += 1
                lm = max(lm, height[l])
                res += lm - height[l]

            }else{
                r -= 1
                rm = max(rm, height[r])
                res += rm - height[r]
            }
    }
    return res
  }
}
