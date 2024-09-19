# OS shell language
Home task for configuration management. To develop an emulator for the OS shell language. It is necessary to make the emulator work as much as possible like a shell session in a UNIX-like OS. The emulator should be run from the real command line, and the file with the virtual file system does not need to be unpacked from the user. The emulator accepts an image of a virtual file system in the form of a file format zip. The emulator should work in CLI mode.

The command line keys are set:
• The username to show in the input prompt.
• The name of the computer to display in the input prompt.
• The path to the archive of the virtual file system.
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