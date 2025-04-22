# Project Euler in Go

This project contains solutions to various Project Euler problems implemented in Go. Each problem is contained within its own directory under the `scripts` folder, allowing for straightforward execution of individual scripts.

## Project Structure

```
project-euler-go
├── scripts
│   ├── problem001
│   │   └── main.go
│   ├── problem002
│   │   └── main.go
│   ├── problem003
│   │   └── main.go
│   └── ... 
├── pkg
│   ├── utils
│   │   └── math.go
│   └── common
│       └── helpers.go
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Getting Started

To get started with this project, ensure you have Go installed on your machine. You can download it from the official Go website.

### Running a Script

To run a specific problem script, navigate to the corresponding directory and execute the `main.go` file. For example, to run Problem 001:

```bash
cd scripts/problem001
go run main.go
```

### Dependencies

This project uses Go modules for dependency management. You can install the required dependencies by running:

```bash
go mod tidy
```

### Makefile

The Makefile provides commands to simplify the execution of scripts and management of dependencies. You can use the following commands:

- `make run-problem001` - Runs the solution for Problem 001.
- `make run-problem002` - Runs the solution for Problem 002.
- `make run-problem003` - Runs the solution for Problem 003.
- Add additional commands as needed for other problems.

## Contributing

Feel free to contribute by adding solutions to more Project Euler problems. Ensure that each solution is placed in its own directory under `scripts` and follows the same structure.

## License

This project is open source and available under the MIT License.