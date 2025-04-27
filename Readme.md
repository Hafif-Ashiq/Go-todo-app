# To-Do Application

This is a simple command-line To-Do application written in Go. It allows users to manage their tasks efficiently by adding, listing, marking as complete, and removing tasks. The tasks are stored persistently in a JSON file.

## Features

- Add a new task
- List all tasks
- Mark a task as complete
- Remove a task

## Project Structure

- `main.go`: The entry point of the application.
- `todo/`: Contains the core logic for managing tasks.
  - `todo.go`: Defines the `Todo` and `TodoList` structures and their methods.
- `commands/`: Contains the command implementations.
  - `add.go`: Command for adding tasks.
  - `list.go`: Command for listing tasks.
  - `update.go`: Command for marking tasks as complete.
  - `delete.go`: Command for removing tasks.
- `tasks.json`: Stores the tasks persistently.

## Usage

### Prerequisites

- Go 1.23.3 or later installed on your system.

### Running the Application

1. Clone the repository.
2. Navigate to the project directory.
3. Run the application using the following command:

   ```bash
   go run main.go <command> [arguments]
   ```

### Commands

- **Add a Task**
  ```bash
  go run main.go add "Task Title"
  ```

- **List All Tasks**
  ```bash
  go run main.go list
  ```

- **Mark a Task as Complete**
  ```bash
  go run main.go mark-complete <task_id>
  ```

- **Remove a Task**
  ```bash
  go run main.go remove <task_id>
  ```

## Example

1. Add a task:
   ```bash
   go run main.go add "Learn Go"
   ```

2. List tasks:
   ```bash
   go run main.go list
   ```

3. Mark a task as complete:
   ```bash
   go run main.go mark-complete 1
   ```

4. Remove a task:
   ```bash
   go run main.go remove 1
   ```

## License

This project is licensed under the MIT License.