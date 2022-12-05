import io
import re

MOVE_RE = r"move ([0-9]+) from ([0-9]?) to ([0-9]?)"

def parse_input(path : str):
    state = []
    moves = []
    with io.open(path, 'r') as file:
        states_ended = False
        for line in file.readlines():
            if line != '\n':
                if(not states_ended):
                    state.append(line)
                else:
                    moves.append(line)
            else:
                states_ended = True
    
    piles = state[-1].split('   ')
    max_size = int(piles[-1].strip())
    
    state_decoded = []
    for i in range(max_size):
        state_decoded.append([])
    
    for i in range(len(state)-2,-1,-1):
        current = state[i]
        for j in range(4,len(current)+1, 4):
            c_s = current[j-4:j].strip()
            if(c_s != ''):
                state_decoded[j//4-1].append(c_s)
                if('][' in c_s):
                    print(state_decoded[0])
                    print(state_decoded[1])
                    print(state_decoded[2])
                    print(state_decoded[3])
                    print(state_decoded[4])
                    print(state_decoded[5])
                    print(state_decoded[6])
                    print(state_decoded[7])
                    print(state_decoded[8])
                    print("Error: \n", current, current.replace('   ',''), element, c_s,'\n----')
    
    moves_decoded = []
    for move in moves:
        match = re.search(MOVE_RE, move)
        if(match):
            count = int(match.group(1))
            from_p = int(match.group(2))
            to_p = int(match.group(3))
            moves_decoded.append((count, from_p, to_p))
        else:
            print("ERROR", move)
            break

    return (state_decoded, moves_decoded)

def operate_crane9000(state, moves):
    for count, from_p, to_p in moves:
        if(len(state[from_p-1]) < 0):
            print('ERROR: ', count, from_p, to_p)
            return

        for i in range(count):
            el = state[from_p-1].pop()
            state[to_p - 1].append(el)

def operate_crane9001(state, moves):
    for count, from_p, to_p in moves:
        if(len(state[from_p-1]) < 0):
            print('ERROR: ', count, from_p, to_p)
            return
        
        for el in state[from_p-1][-1*count:]:    
            state[to_p-1].append(el)

        for i in range(count):    
            state[from_p-1].pop()

if __name__ == "__main__":
    name = "input.txt"
    input = parse_input(name)
    state = input[0]
    moves = input[1]
    
    operate_crane9001(state, moves)

    result = ""
    for s in state:
        if(len(s) == 0):
            result += '-'
        else:
            result += s[-1].strip('[').strip(']')
    print(result)