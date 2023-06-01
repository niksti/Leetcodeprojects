class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        for index, number in enumerate(nums):
            for index2, number2 in enumerate(nums):#[index:],index):
                if index != index2 and number + number2 == target:
                    return [index,index2]
