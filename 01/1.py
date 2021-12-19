
with open("input.txt", "r") as file:
    count = 0
    prev = int(file.readline().strip())

    for line in file:
        next = int(line.strip())
        if next > prev:
            count = count + 1

        prev = next

    print(f"Increases: {count}")
