# Go Password Generator

This is a command-line password generator tool written in Go. It generates random, secure passwords based on your specified criteria.

## Features

- Generates passwords of desired length (minimum 8 characters).
- Includes options for lowercase letters, uppercase letters, numbers, and symbols.
- Customizable symbol set (defaults to common symbols).

## Installation

### Prerequisites

- **Go version 1.17 or later** installed on your system. You can check your Go version by running `go version` in your terminal.
- A text editor or IDE for editing Go code.

### Building

1. **Clone this repository** or download the source code.

2. **Open a terminal** in the directory containing the source code (`password_generator.go`).

3. **Run the following command to build the executable**:
   ```bash
   go build -o password_generator password_generator.go
