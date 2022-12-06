import io

def parse_input(path : str) -> str:
    with io.open(path, 'r') as file:
        for line in file.readlines():
            return line

def find_start_position(package : str, size: int) -> int:
    last_chars = []
    for i in range(len(package)):
        if(len(last_chars) == size):
            return i
        if(package[i] not in last_chars):
            last_chars.append(package[i])
        else:
            el = '*'
            while(el != package[i]):
                el = last_chars.pop(0)
            last_chars.append(package[i])
    return -1


if __name__ == "__main__":
    val = parse_input("input.txt")
    position = find_start_position(val, 4)
    position_2 = find_start_position(val, 14)
    print(position, position_2)