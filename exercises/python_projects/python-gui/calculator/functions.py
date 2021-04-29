from tkinter import *

add = "Add"
sub = "Subtract"
mul = "Multiply"
div = "Divide"

num1: float
operator: str

def button_click(e, number):
    e: Entry
    e.insert(END, number)

def button_clear(e):
    e: Entry
    e.delete(0, END)

def button_add(e):
    e: Entry
    global num1
    global operator

    operator = add
    num1 = float(e.get())
    button_clear(e)

def button_sub(e):
    e: Entry
    global num1
    global operator
    
    operator = sub
    num1 = float(e.get())
    button_clear(e)

def button_mul(e):
    e: Entry
    global num1
    global operator
    
    operator = mul
    num1 = float(e.get())
    button_clear(e)

def button_div(e):
    e: Entry
    global num1
    global operator
    
    operator = div
    num1 = float(e.get())
    button_clear(e)

def button_equals(e):
    e: Entry
    global num1
    global operator
    
    switcher = {
        add: num1 + float(e.get()),
        sub: num1 - float(e.get()),
        mul: num1 * float(e.get()),
        div: num1 / float(e.get())
    }
    
    result = switcher.get(operator, "invalid operator")
    button_clear(e)
    button_click(e, result)
