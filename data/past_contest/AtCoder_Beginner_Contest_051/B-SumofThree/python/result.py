k, s = map(int, input().split(" "))

result = 0

for x in range(k+1):
  for y in range(k+1):
    z = s - x - y
    if z >= 0 and z <= k:
      result += 1

print(result)