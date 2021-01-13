import random as r

words = ["Hello", "lol", "evil", "random"]
correct = []
incorrect = []

word = words[r.randint(0, len(words))]
tries = len(word) + 5

for a in word:
    correct.append('_')

def guess(x):
    if word.__contains__(x):
        for y in x:
            for i in range(len(word)):
                if word[i] == y:
                    correct[i] = y
        return True
    else:
        for y in x:
            incorrect.append(y)
        return False


while True:
    if tries < 1:
        print("\n\nGame over!")
        break

    print(correct, "\nGraveyard: ", incorrect, "\nTries: ", tries)
    g = input("Enter guess: ")
    b = guess(g)
    if correct.__contains__('_') == False:
        print("\n\nYou win! Tries: ", tries)
        break
    tries -= 1
