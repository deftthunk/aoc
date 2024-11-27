alphas = {
    'zero': 0,
    'one': 1,
    'two': 2,
    'three': 3,
    'four': 4,
    'five': 5,
    'six': 6,
    'seven': 7,
    'eight': 8,
    'nine': 9,
}


with open('puzzle1_input.txt', 'r') as fh:
    value = 0
    for line in fh.readlines():
        group = []
        for c in line:
            if c.isdigit():
                group.append(c)

        str_num = ''.join([group[0],group[-1]])
        print(f"str_num: {str_num}")
        real_num = int(str_num)
        value += real_num

    print(value)
