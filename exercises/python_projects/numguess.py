import random as r

def guess(x, y):
    return x == y # Compares the guessed number x to the solution y


val = r.randint(0, 100) # Randomly selects a number val between 0 and 100
tries = 0 # Number of times the player has guessed

# Cycles through the game loop until the player guesses the number val correctly
while True:
    g = input("Guess a number: ")
    tries += 1
    correct = guess(int(g), val)

    if correct:
        print("\nCorrect! " + tries + " attempts.\n")
        break
    else:
        print("\nIncorrect! Guess again.\n")
