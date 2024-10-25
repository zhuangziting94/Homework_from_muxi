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
