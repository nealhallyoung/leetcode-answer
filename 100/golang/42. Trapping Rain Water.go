// 42. Trapping Rain Water
// double pointer methods
func trap(height []int) int {
    sum := 0
    for i:=0;i<len(height)-1;i++{
        maxleft:=0
        for j:=i-1;j>=0;j--{
            if(height[j]>maxleft){
                maxleft=height[j]
            }
        }
        maxright:=0
        for j:=i+1;j<len(height);j++{
            if(height[j]>maxright){
                maxright=height[j]
            }
        }
        minH := min(maxleft,maxright)
        if(height[i]<minH){
            sum=sum+(minH-height[i])
        }
        
    }
    return sum
}
