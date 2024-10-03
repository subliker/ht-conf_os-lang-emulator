# OS shell language
Home task for configuration management. To develop an emulator for the OS shell language. It is necessary to make the emulator work as much as possible like a shell session in a UNIX-like OS. The emulator should be run from the real command line, and the file with the virtual file system does not need to be unpacked from the user. The emulator accepts an image of a virtual file system in the form of a file format zip. The emulator should work in CLI mode.

The command line keys are set:\
• The username to show in the input prompt.\
• The name of the computer to display in the input prompt.\
• The path to the archive of the virtual file system.\
• The path to the start script.

The start script is used for the initial execution of a given list
of commands from a file.

It is necessary to support ls, cd and exit commands in the emulator, as well
as the following commands:
1. echo.
2. whoami.
3. uniq.

All emulator functions must be covered by tests, and
2 tests must be written for each of the supported commands.

### Requirements
![golang](https://badgen.net/static/go/1.22.2/green?icon=github)<br/>
You can install Golang <a href="https://go.dev/doc/install">there</a>

### Run
1. Clone repository
2. In main directory
    ```
    go build -o bin/emu.exe cmd/emu/main.go
    
    //for windows
    ./bin/emu.exe

    //for linux
    ./bin/emu
    ```

### Tests
Run all tests
```
go test -v ./...
```

### Help
Run  ```./emu-x64.exe --help``` to get flags information

### Usage examples
- Using flags
  ```
  /*
      It runs cli with username subliker and pc name mega-pc.
      It uses fs from test.zip. It uses start script test.nya
  */
  ./emu-x64.exe --username subliker --pcname mega-pc --apath test.zip --startpath test.nya
  ```
- Using start script
  ```
  // start script inner
  echo 'It is echo from start script!'
  ```
  ![image](https://github.com/user-attachments/assets/aa5c7c39-e981-415d-aadd-8b0fefa34f29)
  
- Using archive & ```ls```
  ```
  // archive inner
  a/b.txt
  c.txt
  ```
  ![image](https://github.com/user-attachments/assets/365fc2ca-31ae-44fd-8b8e-ffe57c9180c2)

- Using ```cd```\
  \
  ![image](https://github.com/user-attachments/assets/7874fa81-6ce3-4ee6-9cc8-221b3bdad8d5)

- Using ```exit```\
  \
  ![image](https://github.com/user-attachments/assets/63b95f24-b22b-41e6-8c24-470a0707ca95)

- Using ```echo```\
  \
  ![image](https://github.com/user-attachments/assets/50a2b7dc-4e11-46b9-b989-f80ad9d2ea5e)

- Using ```whoami```\
  \
  ![image](https://github.com/user-attachments/assets/e8f46265-d4bf-44e2-8885-247deb0396c4)

- Using ```uniq```
  ```
  // a.txt inner
  1
  1
  1
  2
  2
  1
  ```
  ![image](https://github.com/user-attachments/assets/5a1aa337-925d-4d66-8120-2e8d306f1e99)
