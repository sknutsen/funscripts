name = ""
city = ""
country = ""

story = "\n{} was from {}, in the country of {}\n"

def lazy_libs():
    name = input("Name of the character: ")
    city = input("City the character is from: ")
    country = input("Country the city is in: ")
    print(story.format(name, city, country))
    print("that's all folks!\n")


lazy_libs()
