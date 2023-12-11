class Solution {
    func canPlaceFlowers(_ flowerbed: [Int], _ n: Int) -> Bool {
        if flowerbed.count < n{
            return false
        }
        if flowerbed.count == 0 && n == 0{
			return true
		}
        if flowerbed.count == 0 && n > 0{
			return true
		}
		if flowerbed.count == 1 && flowerbed[0] == 1 && n > 0{
			return false
		}
        if flowerbed.count == 1 && flowerbed[0] == 1 && n == 0 {
			return true
		}
        if flowerbed.count == 1 && flowerbed[0] == 0 && (n == 0 || n == 1) {
			return true
		}
        if flowerbed.count == 1 && flowerbed[0] == 0 && n > 1 {
			return false
		}
        let first = 0
        let last = flowerbed.count-1
        var array = flowerbed
        var number=0
        print(last)
        let con = n
        // 有多少可以种花的空位
        for (k,_) in array.enumerated() {
            // 首部
            if k == first{
                if array[first+1] == 0 && array[first] == 0{
                    array[first] = 1
                    number+=1
                }
            }else if k == last{ // 尾部
                print("尾部")
                print((array[last-1] == 0) && (array[last] == 0))
                print("array[last-1]: \(array[last-1])")
                print("array[last]: \(array[last])")
                if (array[last-1] == 0) && (array[last] == 0){
                    array[last] = 1
                    number+=1
                }
            }else {
                if array[k-1] == 0 && array[k+1] == 0 && array[k] == 0{
                    array[k]=1
                    number+=1
                }
            }
        } // for

        
        if con <= number {
            return true
        }else{
            return false
        }

    }
}