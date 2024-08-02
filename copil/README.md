# copil

A Chainlit application written in Python that allows for the Q and A over a SQLite database containing an inventory of items, and has a widget component that can be embedded in a webpage.

Might try to integrate this copilot into my fork of homebox.

## Usage

To run the chainlit server

    chainlit run app.py

![chatbot](./assets/Chatbot.jpeg)

Serve the index.html file _somehow_ which includes the embedded widget.

    python3 -m http.server 8001

Then navigate to `http://localhost:8001` in your browser.

![chatbot](./assets/Copilot.jpeg)
