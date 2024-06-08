import os
import streamlit as st
from dotenv import load_dotenv

from langchain_community.utilities.sql_database import SQLDatabase
from langchain_community.agent_toolkits.sql.toolkit import SQLDatabaseToolkit
from langchain_community.callbacks import StreamlitCallbackHandler
from langchain_community.chat_message_histories.streamlit import StreamlitChatMessageHistory
from langchain_core.runnables import RunnableConfig
from langchain_openai import ChatOpenAI
from langchain_core.messages import HumanMessage
from langchain.memory import ConversationBufferMemory
from langchain.agents import AgentExecutor, AgentType
from langchain_community.agent_toolkits.sql.base import create_sql_agent
from langchain.tools.render import ToolsRenderer, render_text_description
from langchain.agents import create_react_agent

from prompt import REACT_PROMPT

# Load environment variables
load_dotenv()
required_env_vars = [
    "LANGCHAIN_ENDPOINT", "LANGCHAIN_API_KEY", "LANGCHAIN_PROJECT", "LANGCHAIN_TRACING_V2",
    "OPENAI_API_KEY",
]

# Collect environment variables and check if they are all set.
vals = {var: os.getenv(var) for var in required_env_vars}
valid = all(v is not None for v in vals.values())

# Streamlit page configuration
st.set_page_config(page_title="Inventory Chat", page_icon="📝")
st.title("📝 Inventory Chat")

# Initialize conversation history and memory
msgs = StreamlitChatMessageHistory(key="chat_history")
memory = ConversationBufferMemory(
    chat_memory=msgs, return_messages=True, memory_key="chat_history", output_key="output"
)
llm = ChatOpenAI(model="gpt-4o", temperature=0, streaming=True)
db=SQLDatabase.from_uri("sqlite:///inv.db")
toolkit = SQLDatabaseToolkit(llm=llm, db=db)
tools = toolkit.get_tools()
agent = create_react_agent(llm, tools, REACT_PROMPT)
executor = AgentExecutor(
    agent=agent, # type: ignore
    tools=tools,
    memory=memory,f
    verbose=True,
    handle_parsing_errors=True,
    max_iterations=10
)
session_id = "41734cdb-dab1-4030-a441-497e6a000100"

# Sidebar for environment variables status
with st.sidebar:
    status_text = ":green[Environment is valid]" if valid else ":red[Environment is invalid]"
    with st.expander(status_text, expanded=False):
        on = st.toggle("Show values")
        for var in required_env_vars:
            stat_out = "✅" if vals[var] is not None else "❌"
            st.write(f"{stat_out} {var}")
            if on:
                st.code(f"{vals[var]}")

    # SessionID input
    st.write("Session ID")
    session_id = st.text_input("Enter session ID", session_id)

# Reset chat history if button is clicked
if len(msgs.messages) == 0 or st.sidebar.button("Reset chat history"):
    msgs.clear()
    msgs.add_ai_message("How can I help you?")
    st.session_state.steps = {}


# Display conversation messages
avatars = {"human": "user", "ai": "assistant"}
for idx, msg in enumerate(msgs.messages):
    with st.chat_message(avatars[msg.type]):
        for step in st.session_state.steps.get(str(idx), []):
            if step[0].tool == "_Exception":
                continue
            with st.status(f"**{step[0].tool}**: {step[0].tool_input}", state="complete"):
                st.write(step[0].log)
                st.write(step[1])
        st.write(msg.content)


# Process chat input and generate response
if prompt_str := st.chat_input(placeholder="What categories does the db contain?"):
    st.chat_message("user").write(prompt_str)

    # Display agent's response
    with st.chat_message("assistant"):
        st_cb = StreamlitCallbackHandler(st.container())
        cfg = RunnableConfig(callbacks=[st_cb], metadata={"session_id": session_id})
        response = executor.invoke({"input": prompt_str}, cfg)
        st.write(response["output"])
