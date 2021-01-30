from tkinter import *
import requests
import json

root = Tk()
root.title("API Client")
root.geometry("400x50")

try:
    api_request = requests.get("http://dummy.restapiexample.com/api/v1/employees")
    print(api_request)
    resp = json.loads(api_request.content)
except Exception as e:
    print(e.with_traceback())

label = Label(root, text=resp["data"][0])
label.pack()

root.mainloop()
