# inventory

## Description

This repository contains a few projects related to inventory management in relation to my personal collection of electronics equipment. 

### The Problem

I have a lot of cables, and I forget where they are. 

### The Solution

The goal of this repository is to create an easy way to keep track of what I have, where it is, and what it is used for.

## Projects

### [Go Inventory](./goinv/README.md)

A RESTful API written in Go that allows for the management of an inventory of items.

### [Python Inventory](./pyinv/README.md)

A Streamlit application written in Python that allows for the Q and A over a SQLite database containing an inventory of items.

## Data

The data I'm using for testing is real data from my collection of electronics equipment, and can be found in the [inventory.csv](./inventory.csv) file.

### Example Data

| Category | Item                    | Quantity | Location |
|----------|-------------------------|----------|----------|
| Adapter  | AC to USB-A with button | 1        | w1       |
| Device   | Asus mini laptop        | 1        | w1       |
| Cable    | USB-A to USB-C          | 8        | black    |
| Adapter  | Single 5V 1A            | 7        | w1       |
| Device   | Raspberry Pi 4          | 1        | w2       |
etc ...
