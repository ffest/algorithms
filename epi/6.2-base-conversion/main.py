import functools
import string


def convert_base(num: str, b1: int, b2: int) -> str:
    def construct_from_base(number, base):
        return ("" if number == 0 else
                construct_from_base(number // base, base) + string.hexdigits[number % base].upper())

    is_negative = num[0] == "-"
    num_as_int = functools.reduce(
        lambda x, c: x * b1 + string.hexdigits.index(c.lower()),
        num[is_negative:], 0)
    return ("-" if is_negative else '') + ('0' if num_as_int == 0 else construct_from_base(num_as_int, b2))


print(convert_base("615", 7, 13))
