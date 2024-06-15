Go Password Generator
This is a command-line password generator tool written in Go. It generates random, secure passwords based on your specified criteria.

Features
Generates passwords of desired length (minimum 8 characters).
Includes options for lowercase letters, uppercase letters, numbers, and symbols.
Customizable symbol set (defaults to common symbols).
Installation
Prerequisites:

Go version 1.17 or later installed on your system. You can check your Go version by running go version in your terminal.
A text editor or IDE for editing Go code.
Building:

Clone this repository or download the source code.
Open a terminal in the directory containing the source code (password_generator.go).
Run the following command to build the executable:
go build -o password_generator password_generator.go
This will create an executable file named password_generator in the same directory.

Usage
Run the password_generator executable with desired options:

./password_generator [flags]
Available Flags:

-length: Sets the desired password length (default 16, minimum 8).
-lowercase: Include lowercase letters (default true).
-uppercase: Include uppercase letters (default true).
-numbers: Include numbers (default true).
-symbols: Include symbols (default true).
-symbols="custom_symbols": Use a custom symbol set (replace "custom_symbols" with your desired characters).
Examples:

Generate a 20-character password with all character sets:
./password_generator -length 20
Generate a 12-character password with only numbers and symbols:
./password_generator -length 12 -lowercase=false -uppercase=false
Generate a 15-character password with a custom symbol set:
./password_generator -length 15 -symbols="!@#$%&*()"
