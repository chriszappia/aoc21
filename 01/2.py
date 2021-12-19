
with open("input.txt", "r") as file:
    count = 0
    previous_sum = 0
    prev_2 = int(file.readline().strip())
    prev_1 = int(file.readline().strip())

    for line in file:
        current = int(line.strip())
        current_sum = current + prev_1 + prev_2
        if previous_sum and current_sum > previous_sum:
            count = count + 1
        print(f"C: {current_sum}, P: {previous_sum} inc: {current_sum > previous_sum}")
        prev_2 = prev_1
        prev_1 = current
        previous_sum = current_sum

    print(f"Increases: {count}")
