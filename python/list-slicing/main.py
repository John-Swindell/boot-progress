def get_champion_slices(champions):
    first = champions[2:]
    next = champions[:-1] 
    last = champions[::2]
    return first, next, last

