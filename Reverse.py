class Solution:
    def reverse(self, x: int) -> int:
        result = 0
        neg = False
        if (x < 0):
            x *= (-1)
            neg = True
        while (x != 0):
            y = x % 10
            x = (x - y) / 10
            result = (result + y) * 10
        if (neg):
            result = result / (-10)
        else:
            result = result/10

        if (result < -(2 ** 31) or result >= 2 ** 31):
            return 0
        return int(result)
