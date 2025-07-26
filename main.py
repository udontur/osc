numberOfItems = int(input("Number of items? "))

marks = [
    float(mark)
    for mark in
    input("Enter the marks: ").split()
]
denominators = [
    float(denominators)
    for denominators in
    input("Enter the denominators: ").split()
]
weights = [
    float(weight)
    for weight in
    input("Enter the weights: ").split()
]

totalWeight = float(sum(weights))

exactMark = 0
for i in range(numberOfItems):
    mark = marks[i] / denominators[i]
    weight = weights[i] / totalWeight

    sectionMark = mark * weight
    exactMark += sectionMark

exactMark *= 100
finalMark = round(exactMark)
exactMark = "{:.3f}".format(exactMark)

GREEN = "\033[92m"
RESET = "\033[39m"

print()
print(f"{GREEN}Final Mark: {finalMark}{RESET}")
print(f"Exact Mark: {exactMark}")
