import re

with open("input.txt") as fh:
    for line in fh.readlines():
        game, draws = line.split(':')
        gameNum = re.search(r'\d+', game).group(0)
        for draw in draws.split(';'):
            count, color = 

