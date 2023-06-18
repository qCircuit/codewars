'''
def points(games):
    total_points = 0
    for game in games:
        game =  game.split(":")
        a, b = int(game[0]), int(game[1])
        if a > b:
            total_points +=  3
        elif a == b:
            total_points += 1
        else:
            pass

    return total_points
'''
games = ['1:0','2:0','3:0','4:0','2:1','3:1','4:1','3:2','4:2','4:3']

def points(games):
    return sum((x == y) + 2 * (x > y) for x, y in (game.split(":") for game in games))

print(points(games))