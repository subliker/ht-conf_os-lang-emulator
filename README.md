# OS Shell Language Emulator

This project is a home task for configuration management. 

## Table of Contents

- [Task](#task)
- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [Running Tests](#running-tests)
- [Help](#help)
- [Usage Examples](#usage-examples)

## Task

To develop an emulator for the OS shell language. It is necessary to make the emulator work as much as possible like a shell session in a UNIX-like OS. The emulator should be run from the real command line, and the file with the virtual file system does not need to be unpacked from the user. The emulator accepts an image of a virtual file system in the form of a file format zip. The emulator should work in CLI mode.

### Emulator Commands
- `ls`
- `cd`
- `exit`
- `echo`
- `whoami`
- `uniq`

### Command Line Keys
- **Username**: Displays in the input prompt.
- **Computer Name**: Displays in the input prompt.
- **Path to Virtual File System Archive**: Specifies the ZIP file.
- **Path to Start Script**: Used for initial command execution.

All emulator functions must be covered by tests, and two tests must be written for each of the supported commands.

## Requirements

![golang](https://badgen.net/static/go/1.22.2/green?icon=github)<br/>
You can install Golang [here](https://go.dev/doc/install).

## Getting Started

To get a local copy up and running, follow these steps:

1. **Clone the repository**:

   ```bash
    git clone https://github.com/subliker/ht-conf_os-lang-emulator
    cd ht-conf_os-lang-emulator
    ```

3. **Build the emulator**:

   In the main directory, run:

    ```bash
    go build -o bin/emu.exe cmd/emu/main.go
    ```

4. **Run the emulator**:

   - For Windows:
     ```
     ./bin/emu.exe
     ```

   - For Linux:
     ```
     ./bin/emu
     ```

## Running Tests

To run all tests, execute:

```
go test -v ./...
```

## Help

For flags information, run:

```
./emu-x64.exe --help
```

## Usage Examples

- **Using Flags**:

  ```
  ./emu-x64.exe --username subliker --pcname mega-pc --apath test.zip --startpath test.nya
  *This runs the CLI with the username `subliker`, PC name `mega-pc`, using the file system from `test.zip` and the start script `test.nya`.*
  ```

- **Using Start Script**:

  ```
  // Start script content
   echo 'It is echo from start script!'
  ```
  
  ![Start Script Example](https://github.com/user-attachments/assets/aa5c7c39-e981-415d-aadd-8b0fefa34f29)

- **Using Archive & `ls`**:
  ```
  // Archive content
  a/b.txt
  c.txt
  ```
  
  ![List Files Example](https://github.com/user-attachments/assets/365fc2ca-31ae-44fd-8b8e-ffe57c9180c2)

- **Using `cd`**:
  
  ![Change Directory Example](https://github.com/user-attachments/assets/7874fa81-6ce3-4ee6-9cc8-221b3bdad8d5)

- **Using `exit`**:
  
  ![Exit Example](https://github.com/user-attachments/assets/63b95f24-b22b-41e6-8c24-470a0707ca95)

- **Using `echo`**:
  
  ![Echo Example](https://github.com/user-attachments/assets/50a2b7dc-4e11-46b9-b989-f80ad9d2ea5e)

- **Using `whoami`**:
  
  ![Who Am I Example](https://github.com/user-attachments/assets/e8f46265-d4bf-44e2-8885-247deb0396c4)

- **Using `uniq`**:
  ```
  // a.txt content
  1
  1
  1
  2
  2
  1
  ```
  
  ![Unique Example](https://github.com/user-attachments/assets/5a1aa337-925d-4d66-8120-2e8d306f1e99)
