import functools


def rabin_karp(t: str, s: str) -> int:
    if len(s) > len(t):
        return -1

    base = 26
    t_hash = functools.reduce(lambda h, c: h * base + ord(c), t[:len(s)], 0)
    s_hash = functools.reduce(lambda h, c: h * base + ord(c), s, 0)
    power_s = base ** max(len(s)-1, 0)

    for i in range(len(s), len(t)):
        if t_hash == s_hash and t[i-len(s):i] == s:
            return i-len(s)
        t_hash -= ord(t[i-len(s)])*power_s
        t_hash = t_hash * base + ord(t[i])

    if t_hash == s_hash and t[-len(s)] == s:
        return len(t)-len(s)
    return -1


s = "pes"
t = "pesvishelpogulat"
print(rabin_karp(t, s))
