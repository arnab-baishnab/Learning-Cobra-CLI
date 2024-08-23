# Learning-Cobra-CLI

## Todo CLI

A command-line interface (CLI) application for managing todos, written in Go using the Cobra library.

## Features

- Add new todos
- List all todos
- Mark todos as completed
- Delete specific todos
- Clear all todos

## Prerequisites

- Go 1.18 or later
- [Cobra CLI](https://github.com/spf13/cobra)

## Installation

1. **Clone the Repository**

   ```sh
   git clone https://github.com/arnab-baishnab/Learning-Cobra-CLI.git
   cd Learning-Cobra-CLI
   ```
2. **Build the CLI Application**
   ```sh
   go build -o todo
   ```
3. **Run the CLI Application**
   ```sh
   ./todo [command] [flags]
   ```
## Usage

### Adding a New Todo

To add a new todo, use the `add` command followed by the task description:

```sh
./todo add "Buy groceries"
```

### Listing All Todos

To list all todos, use the list command:

```sh
./todo list
```
### Completing a Todo

To mark a todo as completed, use the complete command followed by the index of the todo:

```sh
./todo complete 1
```

### Deleting a Todo

To delete a specific todo, use the delete command followed by the index of the todo:

```sh
./todo delete 1
```


### Clearing All Todos

To delete all todos, use the clear command:

```sh
./todo clear
```


## Commands

 - `add [task]`
Adds a new todo item.

- `complete [index]`
Marks the todo item at the specified index as completed.

- `delete [index]`
Deletes the todo item at the specified index.

- `clear`
Deletes all todo items.

- `list`
Lists all todo items.


## Error Handling

If you encounter any errors, such as invalid index or file access issues, the CLI will provide appropriate error messages.

## Contributing

If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch:
   ```sh
   git checkout -b feature-branch
   ```
3. Make your changes.
4. Commit your changes:
   ```sh 
   git commit -m "Add new feature"
   ```
5. Push to the branch:
   ```
   git push origin feature-branch
   ```
6. Create a new Pull Request


## Acknowledgments
   - Cobra CLI for creating the CLI framework
   - Go for the programming language
