class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == "0" or num2 == "0":
            return "0"
        result = [0] * (len(num1)+len(num2))
        for i in reversed(range(len(num1))):
            for j in reversed(range(len(num2))):
                tmp = int(num2[j])*int(num1[i]) + result[i+j+1]
                result[i+j+1] = tmp % 10
                result[i+j] += tmp//10
        zeros = 0
        while result[zeros] == 0:
            zeros += 1

        return ''.join(map(str, result[zeros:]))


solution = Solution()
print(solution.multiply("9", "99"))
