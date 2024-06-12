import streamlit as st
import sqlite3
import pandas as pd
from streamlit_extras.dataframe_explorer import dataframe_explorer
from streamlit_tags import st_tags

# Set up database connection
connection = sqlite3.connect('inventory.db')
cursor = connection.cursor()

# Get unique categories and locations from the database
categories_query = "SELECT DISTINCT category FROM items"
categories = pd.read_sql(categories_query, connection)['category'].tolist()

locations_query = "SELECT DISTINCT location FROM items"
locations = pd.read_sql(locations_query, connection)['location'].tolist()
query = "SELECT * FROM items"

# Get all items from the database
items = pd.read_sql(query, connection)

st.title("Inventory")

# Select an item
st.selectbox("Select an item", items['name'], index=None, placeholder="Choose an item")

with st.expander("Create a new item", expanded=False):
    with st.form(key='new_item_form'):
        name = st.text_input("Name")
        category = st.selectbox("Category", categories)
        location = st.selectbox("Location", locations)
        qty = st.number_input("Quantity", value=1)
        submit_button = st.form_submit_button(label='Submit')

        if submit_button:
            cursor.execute("INSERT INTO items (name, category, location, qty) VALUES (?, ?, ?, ?)", (name, category, location, qty))
            connection.commit()
            st.write(f"Item {name} added to the inventory")
            st.experimental_rerun()

# 3 column layout of containers, one container for each storage location
cols = st.columns(3)

# Display items in each location
col_ind = 0
for location in locations:
    location_items = items[items['location'] == location]
    with cols[col_ind]:
        st.subheader(location)
        st.dataframe(location_items[['name', 'qty']], hide_index=True)
    col_ind += 1
    if col_ind == 3:
        col_ind = 0

