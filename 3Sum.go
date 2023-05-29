func threeSumClosest(nums []int, target int) int {
    if len(nums) == 3{return nums[0]+ nums[1]+nums[2]}
    current := nums[0]
    best := 1000; 
    nums = nums[1:len(nums)]
    //newtarget := target - current
    for index,a := range(nums){
       // newtarget = target - a
        for i := index + 1; i < len(nums);i++{
            if absDiffInt(current + a + nums[i],target) < absDiffInt(best , target) {
                best = current + a + nums[i]
                if best == target {
                return best
                }
            } 
        }
    }
    //if absDiffInt(best, target) > absDiffInt(threeSumClosest(nums,target),target){
        fff := threeSumClosest(nums,target)
        if absDiffInt(best,target) > absDiffInt(fff,target){
            return fff
        }    
    return best
}


func absDiffInt(x, y int) int {
   if x < y {
      return y - x
   }
   return x - y
}
