# HTTP to Postman Converter

## Overview

The `HTTP to Postman Converter` is a command-line tool written in Go that reads an `.HTTP` file and converts it into a Postman collection. This tool allows developers to quickly generate Postman collections from their existing HTTP request files, making it easier to test APIs and share request configurations.

## Features

- Converts HTTP request files (`.HTTP`) into Postman collections.
- Outputs the Postman collection in JSON format.
- Command-line interface with customizable input and output paths.
- Help documentation available through a simple flag.

## Installation

### Prerequisites

- Go 1.16+ installed on your machine.
- Basic knowledge of Go and command-line usage.

### Clone the Repository

```bash
git clone https://github.com/youssef-elassri/http-to-postman.git
cd http-to-postman
```

### Build the Project

```bash
go build -o http-to-postman
```

This will generate an executable file named `http-to-postman`.

## Usage

### Command-Line Arguments

The tool accepts the following command-line arguments:

- `-i`: Path to the input `.HTTP` file. **(Required)**
- `-o`: Path to the output Postman collection file (JSON format). Default is `postman_collection.json`.
- `-h`: Displays the help message.

### Examples

#### Convert an HTTP File to a Postman Collection

```bash
./http-to-postman -http=requests.http -output=my_collection.json
```

This command reads the `requests.http` file and generates a Postman collection saved as `my_collection.json`.

#### Display Help

```bash
./http-to-postman -help
```

This command displays usage information and available flags.

## Project Structure

```
├── LICENSE
├── Readme.MD
├── go.mod
├── main.go
└── pkg
    ├── converter
    │   └── converter.go
    ├── exporter
    │   └── exporter.go
    ├── models
    │   ├── http_request.go
    │   └── postman.go
    ├── parser
    │   └── parser.go
    └── utils
        └── files.go
```

### Description of Directories and Files

- **`LICENSE`**: Contains the licensing terms under which this project is distributed.

- **`Readme.MD`**: This documentation file provides an overview of the project, usage instructions, and contribution guidelines.

- **`go.mod`**: Go module file that specifies the module path and dependencies required for the project.

- **`main.go`**: The main entry point for the application. This file handles command-line input and controls the overall flow of the program.

- **`pkg/`**: Contains the core logic of the application, broken down into sub-packages:

    - **`converter/`**:
        - `converter.go`: Handles the conversion logic, transforming parsed HTTP requests into Postman collections.

    - **`exporter/`**:
        - `exporter.go`: Manages the exporting of Postman collections to JSON files, using the data structures defined in `models/postman.go`.

    - **`models/`**:
        - `http_request.go`: Defines the Go structs for representing HTTP requests, which are the input to the conversion process.
        - `postman.go`: Defines the Go structs for representing Postman collections, following Postman's official JSON schema.

    - **`parser/`**:
        - `parser.go`: Contains the logic to parse `.HTTP` files and extract HTTP requests. This module transforms raw text into structured data models defined in `models/http_request.go`.

    - **`utils/`**:
        - `files.go`: Utility functions for file operations, such as reading HTTP files from disk and writing output files.


## How It Works

1. **Input Parsing**: The tool reads the input `.HTTP` file provided via the `-http` flag.
2. **Collection Generation**: It parses the HTTP requests and generates a corresponding Postman collection.
3. **Output**: The generated Postman collection is written to the specified output file.

## Contributing

If you'd like to contribute to this project, please fork the repository and create a pull request. Feel free to open issues for bug reports or feature requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

For any questions or support, please contact [youssef3.elassri@gmail.com].

---