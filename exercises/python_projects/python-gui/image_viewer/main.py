from tkinter import *
from tkinter import filedialog
from PIL import ImageTk, Image

root: Tk = Tk()
root.title("Images")

my_img = ImageTk.PhotoImage(Image.open("image_viewer/Emil.jpg"))

imgs = [my_img]

img_no: int
my_label: Label
status: Label

img_no = 0
my_label = Label(image=imgs[img_no])
my_label.grid(row=0, column=0, columnspan=3)

status = Label(root, text="Image " + str(img_no + 1) + " of " + str(len(imgs)), bd=1, relief=SUNKEN, anchor=E)

def back():
    global img_no

    img_no -= 1

    set_img()
    check_back()

def check_back():
    global button_back
    global img_no

    if img_no <= 0:
        if button_back["state"] != "disabled":
            button_back["state"] = "disabled"
        return

    if button_next["state"] == "disabled":
        button_next["state"] = "normal"

def next():
    global img_no

    img_no += 1

    set_img()
    check_next()

def check_next():
    global button_next
    global img_no
    global imgs

    if img_no >= len(imgs) - 1:
        if button_next["state"] != "disabled":
            button_next["state"] = "disabled"
        return

    if button_back["state"] == "disabled":
        button_back["state"] = "normal"

def set_img():
    global img_no
    global imgs
    global my_label
    global status

    status["text"] = "Image " + str(img_no + 1) + " of " + str(len(imgs))

    my_label["image"] = imgs[img_no]

# Broken
def add_img():
    global imgs

    filename = filedialog.askopenfilename(initialdir="./", title="Select a file", filetypes=(("png files", "*.png"),("all files", "*.*")))
    new_image = ImageTk.PhotoImage(Image.open(filename))

    imgs.insert(len(imgs), new_image)
    set_img()
    check_back()
    check_next()

back_sate: str = "normal"
next_sate: str = "normal"

if img_no == 0:
    back_sate = "disabled"
if img_no == len(imgs) - 1:
    next_sate = "disabled"

button_back = Button(root, text="<<", command=back, state=back_sate)
button_next = Button(root, text=">>", command=next, state=next_sate)
button_exit = Button(root, text=root.geometry, command=root.quit)

button_add_img = Button(root, text="Add image", command=add_img)

button_add_img.grid(row=1, column=0, columnspan=3)

button_back.grid(row=2, column=0)
button_next.grid(row=2, column=2)
button_exit.grid(row=2, column=1)

status.grid(row=3, column=0, columnspan=3, sticky=W+E)

root.mainloop()
