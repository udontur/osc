#!/usr/bin/env python3

GREEN = "\033[92m"
RESET = "\033[39m"
GRAY = "\033[90m"

def getInput(name):
    print(f"Enter the {name}: ", end="")
    arr = list(map(
        int, input().split()
    ))
    return arr

marks = getInput("marks")
denominators = getInput("denominators")
weights = getInput("weights")

numberOfItems = len(marks)
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

print()
print(f"{GREEN}Final Mark: {finalMark}{RESET}")
print(f"Exact Mark: {exactMark}")
