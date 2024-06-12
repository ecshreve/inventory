
import streamlit as st
import pandas as pd
import sqlite3
import streamlit_pandas as sp

# Set up database connection
connection = sqlite3.connect('inventory.db')
cursor = connection.cursor()

def increase_qty():
    st.session_state.selected_item['qty'] += 1
    st.write(st.session_state.selected_item)
    cursor.execute("UPDATE items SET qty = ? WHERE id = ?", (st.session_state.selected_item['qty'].values[0], st.session_state.selected_item['id'].values[0]))
    connection.commit()

st.title("Inventory Explorer")

if 'selected_item' not in st.session_state:
    st.session_state.selected_item = None

if st.session_state.selected_item is not None:
    st.write("Selected Item")
    st.write(st.session_state)
    st.write(st.session_state.selected_item)
    st.metric("Quantity", st.session_state.selected_item[['qty']])
    st.button("Increase Quantity", key="increase_qty", on_click=increase_qty)

df = pd.read_sql("SELECT * FROM items", connection)

create_data = {
    "category": "multiselect",
    "location": "multiselect",
}

all_widgets = sp.create_widgets(df, create_data, ignore_columns=["id", "qty", "name"])
res = sp.filter_df(df, all_widgets)

st.header("Filtered DataFrame")
event = st.dataframe(
    res, 
    column_order=['qty', 'name', 'category', 'location'], 
    use_container_width=True, 
    hide_index=True,
    on_select='rerun',
    selection_mode=['single-row']
)

if event:
    st.session_state.selected_item = df.iloc[event["selection"]["rows"]] # type: ignore
            

