def str_to_int(s: str) -> int:
    running_sum = 0
    for i in range(1, len(s)):
        running_sum = running_sum * 10 + ord(s[i]) - ord('0')
    if s[0] == '-':
        running_sum *= -1
    return running_sum


def int_to_str(x: int) -> str:
    negative = False
    if x < 0:
        negative = True
        x *= -1

    result = ""
    while x > 0:
        result += str(x % 10)
        x //= 10
    result = result[::-1]
    if negative:
        result = "-" + result

    return result


print(str_to_int("-321"))
print(int_to_str(-321))
