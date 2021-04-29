import random as r

words = ["Hello", "word"] # List of available words
correct = [] # Obsfuscated representation of the correct word
incorrect = [] # List of incorrect guesses

word = words[r.randint(0, len(words))] # Randomly selects a word from the list of words
attempts = 5 # Sets the number of attempts

for a in word:
    correct.append('_') # adds an underscore in place of every character of the word

# Function which determins whether or not a guess x is correct, and is contained in the word
def guess(x):
    if word.lower().__contains__(x.lower()):
        for y in x:
            for i in range(len(word)):
                if word[i].lower() == y.lower():
                    correct[i] = y # Replaces underscore with correct guessed character y of input x iff x is contained in the word
        return True
    else:
        for y in x:
            incorrect.append(y) # Adds character y of guess x to list of incorrect guesses if guess was incorrect
        return False


# Interates through the game loop until number of attempts reach 0
# or the player guesses the word correctly
while True:
    if attempts < 1:
        print("\n\nGame over!") # Ends the game if number of remaining attempts go below 1(i.e hit 0)
        break

    print("\n\n\n", correct, "\nGraveyard: ", incorrect, "\nAttempts: ", attempts) # Prints out the current status of the game at the beginning of each round
    g = input("Enter guess: ")
    b = guess(g)
    if b == False:
        attempts -= 1 # Decreases the number of available attempts by 1 if guess is incorrect
    elif correct.__contains__('_') == False:
        print("\n\nYou win! Attempts ramaining: ", attempts) # Ends the game if word is guessed correctly
        break
