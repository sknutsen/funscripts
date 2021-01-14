import random as r

dice = int(input("Number of dice: ")) # Number of dice to be used
rolls = int(input("Number of rolls: ")) # Number of times to roll the dice

# Cycle through the number of rolls printing out the result of each round 
# in the format of: Roll i - [d1] [d2] ... [dn]
for i in range(rolls):
    output = "Roll {} - ".format(i + 1)
    for j in range(dice):
        output += "[{}] ".format(r.randint(1, 6))
    
    print(output)
