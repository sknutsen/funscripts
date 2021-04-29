from tkinter import *
import sqlite3

root = Tk()
root.title("CRUD app")
#root.geometry("400x400")

global connstring
connstring: str = "addressbook.db"

def run_command(sql):
    # Open connection to database
    conn = sqlite3.connect(connstring)
    c = conn.cursor()

    if type(sql) == tuple:
        sql:tuple
        c.execute(sql[0], sql[1])
    else:
        c.execute(sql)

    # Commit changes and close connection
    conn.commit()
    conn.close()

'''sql = """CREATE TABLE addresses (
    first_name text,
    last_name text,
    address text,
    city text,
    zipcode integer
)"""
run_command(sql)'''

# Create a new record in the database
def submit():
    sql =("INSERT INTO address(first_name, last_name, address, city, zipcode) VALUES(:fname, :lname, :address, :city, :zip)",
        {
            'fname': str(fname.get()),
            'lname': str(lname.get()),
            'address': str(address.get()),
            'city': str(city.get()),
            'zip': int(zip.get())
        }
    )

    run_command(sql)

    fname.delete(0, END)
    lname.delete(0, END)
    address.delete(0, END)
    city.delete(0, END)
    zip.delete(0, END)

    get_all()

# Delete a record in the database
def delete(id):
    sql = "DELETE FROM address WHERE oid=" + str(id)

    run_command(sql)
    get_all()

def update(id):
    global fname_e
    global lname_e
    global address_e
    global city_e
    global zip_e

    sql = ("""UPDATE address SET
        first_name = :first,
        last_name = :last,
        address = :address,
        city = :city,
        zipcode = :zip

        WHERE oid = :oid""",
        {
            'first': fname_e.get(),
            'last': lname_e.get(),
            'address': address_e.get(),
            'city': city_e.get(),
            'zip': zip_e.get(),
            'oid': id
        }
    )

    run_command(sql)

    fname_e.delete(0, END)
    lname_e.delete(0, END)
    address_e.delete(0, END)
    city_e.delete(0, END)
    zip_e.delete(0, END)

    editor.destroy()

    get_all()

# Update a record in the database
def edit(id):
    record = get_one(id)

    global editor
    editor = Tk()
    editor.title("Edit record")

    # Input labels
    fname_l = Label(editor, text="given name")
    fname_l.grid(row=0, column=0, padx=20)

    lname_l = Label(editor, text="surname")
    lname_l.grid(row=1, column=0, padx=20)

    address_l = Label(editor, text="address")
    address_l.grid(row=2, column=0, padx=20)

    city_l = Label(editor, text="city")
    city_l.grid(row=3, column=0, padx=20)

    zip_l = Label(editor, text="zipcode")
    zip_l.grid(row=4, column=0, padx=20)

    # Globals
    global fname_e
    global lname_e
    global address_e
    global city_e
    global zip_e

    # Input fields
    fname_e = Entry(editor, width=30)
    fname_e.insert(0, record[0])
    fname_e.grid(row=0, column=1, padx=20)

    lname_e = Entry(editor, width=30)
    lname_e.insert(0, record[1])
    lname_e.grid(row=1, column=1, padx=20)

    address_e = Entry(editor, width=30)
    address_e.insert(0, record[2])
    address_e.grid(row=2, column=1, padx=20)

    city_e = Entry(editor, width=30)
    city_e.insert(0, record[3])
    city_e.grid(row=3, column=1, padx=20)

    zip_e = Entry(editor, width=30)
    zip_e.insert(0, record[4])
    zip_e.grid(row=4, column=1, padx=20)

    # Buttons
    submit_btn_e = Button(editor, text="Save record", command=lambda: update(id))
    submit_btn_e.grid(row=5, column=0, columnspan=2, padx=10, pady=10, ipadx=100)

def get_one(id):
    global connstring

    conn = sqlite3.connect(connstring)
    c = conn.cursor()

    c.execute("SELECT *, oid FROM address WHERE oid=" + str(id))
    record = c.fetchone()
    return record

# Get all the records in the database
def get_all():
    global connstring
    global result_frame

    conn = sqlite3.connect(connstring)
    c = conn.cursor()

    c.execute("SELECT *, oid FROM address")
    records = c.fetchall()
    print(records)

    try:
        result_frame.grid_forget()
    except:
        pass

    result_frame = Frame(root)
    result_frame.grid(row=7, column=0, columnspan=2, sticky="nsew")

    for rec in records:
        name_label = Label(result_frame, text=str(rec[5]) + " - " + str(rec[0]) + " " + str(rec[1]))
        name_label.grid(row=int(rec[5]) - 1, column=0)

        upd_btn = Button(result_frame, text="Edit", command=lambda: edit(int(rec[5])))
        upd_btn.grid(row=int(rec[5]) - 1, column=1)

        del_btn = Button(result_frame, text="Delete", command=lambda: delete(int(rec[5])))
        del_btn.grid(row=int(rec[5]) - 1, column=2)

    # Commit changes and close connection
    conn.commit()
    conn.close()

# Frames
input_frame = Frame(root)
input_frame.grid(row=0, column=0, columnspan=2, sticky="nsew")
result_frame = Frame(root)
result_frame.grid(row=7, column=0, columnspan=2, sticky="nsew")

# Input labels
fname_label = Label(input_frame, text="given name")
fname_label.grid(row=0, column=0, padx=20)

lname_label = Label(input_frame, text="surname")
lname_label.grid(row=1, column=0, padx=20)

address_label = Label(input_frame, text="address")
address_label.grid(row=2, column=0, padx=20)

city_label = Label(input_frame, text="city")
city_label.grid(row=3, column=0, padx=20)

zip_label = Label(input_frame, text="zipcode")
zip_label.grid(row=4, column=0, padx=20)

# Input fields
fname = Entry(input_frame, width=30)
fname.grid(row=0, column=1, padx=20)

lname = Entry(input_frame, width=30)
lname.grid(row=1, column=1, padx=20)

address = Entry(input_frame, width=30)
address.grid(row=2, column=1, padx=20)

city = Entry(input_frame, width=30)
city.grid(row=3, column=1, padx=20)

zip = Entry(input_frame, width=30)
zip.grid(row=4, column=1, padx=20)

# Buttons
submit_btn = Button(input_frame, text="Add record", command=submit)
submit_btn.grid(row=5, column=0, columnspan=2, padx=10, pady=10, ipadx=100)
query_btn = Button(input_frame, text="Get data", command=get_all)
query_btn.grid(row=6, column=0, columnspan=2, padx=10, pady=10, ipadx=137)

root.mainloop()
