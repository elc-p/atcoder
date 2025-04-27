s = str(input().strip())
n = len(s)

result_list = []

for bit in range(1 << (n - 1)):
    pattern = ""
    for i in range(n):
        pattern += s[i]
        if i < n - 1 and (bit & (1 << i)):
            pattern += "|"
    result_list += pattern.split("|")

num_list = map(int, result_list)
print(sum(num_list))