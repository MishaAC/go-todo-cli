# Go Todo CLI

A simple **Todo application built with Go**, designed as a command-line interface (CLI).

This project was created after completing the **basic fundamentals section of a Go bootcamp**, with the goal of applying core language concepts in a practical, real-world style project.

---

## ✨ Features

* Interactive terminal menu
* Add new tasks
* List all tasks
* Mark tasks as completed
* Delete tasks
* Input handling that supports spaces and invalid values

---

## 🧠 Concepts Applied

This project focuses on **core Go fundamentals**, including:

* Packages and project structure
* Structs and methods
* Slices and basic data management
* Error handling
* Pointers and constructors (`New` pattern)
* Encapsulation via a task manager
* CLI input handling with `bufio.Reader`

---

## 📁 Project Structure

```
go-todo-cli/
├── main.go            
├── manager/           
│   └── manager.go
├── task/              
│   └── task.go
└── README.md
```

---

## 🚀 Getting Started

### Prerequisites

* Go 1.22 or higher

### Run the application

```bash
go run main.go
```

---

## 🖥️ CLI Preview

```
======================
       TODO CLI
======================
1) Add task
2) List tasks
3) Complete task
4) Delete task
5) Exit

Choose an option:
```

---

## 📝 Notes

* Tasks are stored **in memory only** (no persistence yet).
* This project prioritizes clarity and learning over advanced abstractions.
* Future improvements may include file persistence, flags support, or tests.

---

## 👤 Author

**Cristhian**

Built as part of a Go learning journey 🚀
