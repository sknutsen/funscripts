import random as r

dice = int(input("Number of dice: "))
rolls = int(input("Number of rolls: "))

for i in range(rolls):
    output = "Roll {} - ".format(i + 1)
    for j in range(dice):
        output += "[{}] ".format(r.randint(1, 6))
    
    print(output)
