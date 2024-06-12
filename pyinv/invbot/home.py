import streamlit as st
import pandas as pd
import sqlite3

# Set up database connection
connection = sqlite3.connect('inventory.db')
cursor = connection.cursor()

st.title("Location Explorer")

df = pd.read_sql("SELECT * FROM items", connection)
idx = 0
cols = st.columns(2)
locations = df['location'].unique()

for location in locations:
    location_categories = df[df['location'] == location]['category'].unique()

    with cols[idx%2].container(border=True):
        st.code(f"{location}")

        for category in location_categories:
            items = df[(df['location'] == location) & (df['category'] == category)]
            with st.expander(f"{category} -- {len(items)} -- \t\t{sum(items['qty'])}", expanded=False):
                st.dataframe(items[['name', 'qty']], hide_index=True, use_container_width=True)
    idx += 1