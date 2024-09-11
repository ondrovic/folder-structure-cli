![License](https://img.shields.io/badge/license-MIT-blue)
[![testing](https://github.com/ondrovic/folder-structure-cli/actions/workflows/testing.yml/badge.svg)](https://github.com/ondrovic/folder-structure-cli/actions/workflows/testing.yml)
[![releaser](https://github.com/ondrovic/folder-structure-cli/actions/workflows/releaser.yml/badge.svg)](https://github.com/ondrovic/folder-structure-cli/actions/workflows/releaser.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ondrovic/folder-structure-cli)](https://goreportcard.com/report/github.com/ondrovic/folder-structure-cli)
# Folder Structure CLI

## Overview

Folder Structure CLI is a command-line tool written in Go that creates a folder structure based on a JSON file input. This tool is useful for quickly setting up project structures, organizing files, or replicating directory layouts.

## Features

- Create folders and files based on a JSON structure
- Simple command-line interface
- Recursive creation of nested structures
- Error handling and reporting

## Installation

### Prerequisites

- Go 1.23.1 or later

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/ondrovicfolder-structure-cli.git
   ```
2. Navigate to the project directory:
   ```bash
   cd folder-structure-cli
   ```
3. Build the application:
   ```bash
   go build
   ```

## Usage

### Basic Command
```bash
./folder-structure-cli create [json_file_path] [output_path]
```
- `[json_file_path]`: Path to the JSON file describing the folder structure
- `[output_path]`: Path where the folder structure will be created

### Example

```bash
./folder-structure-cli create structure.json ./output
```

### Version Information

To display the version of the CLI:

```bash
./folder-structure-cli -v
```
or
```bash
./folder-structure-cli --version
```
## JSON Structure

The JSON file should describe the folder structure. Use `null` for files and nested objects for folders.

Example `structure.json`:
```json
{
  "folder1": {
    "subfolder1": {
      "file1.txt": null,
      "file2.txt": null
    },
    "subfolder2": {}
  },
  "folder2": {
    "file3.txt": null
  },
  "file4.txt": null
}
```
This will create:
```
output/
├── folder1/
│   ├── subfolder1/
│   │   ├── file1.txt
│   │   └── file2.txt
│   └── subfolder2/
├── folder2/
│   └── file3.txt
└── file4.txt
```
## Error Handling

The CLI will print error messages for:
- Invalid JSON files
- File read/write errors
- Invalid folder structures

## Development

### Project Structure
```
folder-structure-cli/
├── cmd/
│   ├── root.go
│   └── create.go
├── main.go
└── go.mod
```
- `main.go`: Entry point of the application
- `cmd/root.go`: Defines the root command and version flag
- `cmd/create.go`: Implements the `create` command

### Adding New Commands

To add a new command:

1. Create a new file in the `cmd/` directory (e.g., `cmd/newcommand.go`)
2. Define the command structure and functionality
3. Add the command to the root command in the `init()` function

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
