class Solution {
  func kthSmallestPrimeFraction(_ arr: [Int], _ k: Int) -> [Int] {
    let n = arr.count
    var left = 0.0
    var right = 1.0
    while left < right {
      let mid = (left + right) / 2.0
      var count = 0
      var max = [0, 1]
      var j = 1
      for i in 0..<n {
        while j < n && Double(arr[i]) >= mid * Double(arr[j]) {
          j += 1
        }
        count += n - j
        if j < n && max[0] * arr[j] < arr[i] * max[1] {
          max = [arr[i], arr[j]]
        }
      }
      if count == k {
        return max
      } else if count < k {
        left = mid
      } else {
        right = mid
      }
    }
    return []
  }
}