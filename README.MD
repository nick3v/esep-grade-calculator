# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

## Runtime requirements:
Golang 1.24+
Official releases can be found on go.dev/dl

## Project Structure

- **go.mod** – Defines the module and specifies the required Go version (1.24.4)
- **grade_calculator.go** – Contains the core grade calculation logic, including weighted average computation and numerical-to-letter grade conversion
- **grade_calculator_test.go** – Comprehensive test suite that validates grade calculations by adding sample grades and verifying expected letter grade outputs
- **README.md** – Project documentation and testing instructions

## Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

