# inventory

## Description

This repository contains a few projects related to inventory management, specifically for electronics equipment. Before I found and forked [homebox](https://github.com/ecshreve/homebox-dev) I was using this repo to experiment with different ways to manage my inventory.

### The Problem

I have a lot of cables, and I forget where they are.

### The Solution

The goal of this repository is to create an easy way to keep track of what I have, where it is, and what it is used for.

## Projects

### [Go Inventory](./goinv/README.md)

A RESTful API written in Go that allows for the management of my inventory of items.

### [Python Inventory](./pyinv/README.md)

A Streamlit application written in Python that allows for the Q and A over a SQLite database containing an inventory of items.

### [Copilot Inventory](./copil/README.md)

A Chainlit application written in Python that allows for the Q and A over a SQLite database containing an inventory of items, and has a widget component that can be embedded in a webpage.

Might try to integrate this copilot into my fork of homebox.

## Data

The data I'm using for testing is real data from my collection of electronics equipment, and is stored in a SQLite database.

### Example Data Structure

| Category | Item                    | Quantity | Location |
|----------|-------------------------|----------|----------|
| adapter  | AC to USB-A with button | 1        | box1     |
| device   | Asus mini laptop        | 1        | box1     |
| cable    | USB-A to USB-C          | 8        | shelf1   |
| adapter  | Single 5V 1A            | 7        | w1       |
| device   | Raspberry Pi 4          | 1        | w2       |
etc ...
