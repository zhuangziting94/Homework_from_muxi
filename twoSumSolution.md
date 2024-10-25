# 两数之和的题解  

**思路和算法**  

- 遍历数组寻找`target - nums[i]`,由于`nums[1]`之前的数已经被匹配，所以只要在其之后的
数中寻找即可。

**代码**
```Go
func twoSum(nums []int, target int) []int {
    for i := 0; i< len(nums); i ++{
        for k := i + 1; k < len(nums); k ++{
            if nums[i] + nums[k] == target {
                return []int{i, k}
            }
        }
    }
    return nil
}
```
