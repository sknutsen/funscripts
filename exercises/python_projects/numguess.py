import random as r

def guess(x, y):
    return x == y


val = r.randint(0, 100)
tries = 0

while True:
    g = input("Guess a number: ")
    tries += 1
    correct = guess(int(g), val)

    if correct:
        print("\nCorrect! " + tries + " attempts.\n")
        break
    else:
        print("\nIncorrect! Guess again.\n")
