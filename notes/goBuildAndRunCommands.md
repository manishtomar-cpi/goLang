# Go Build and Run Commands

## 1. Build an Executable File
- Command:  
  go build "file path"  
- This command compiles the Go source code and creates an executable file (.exe).  
- The executable is directly in machine code, so it can run without needing Go installed.

Example:
go build main.go  
This will create an executable file named main.exe (on Windows) or main (on Linux/Mac).

---

## 2. Run the Code Directly
- Command:  
  go run "file path"  
- This command compiles and runs the code immediately, but does not create an executable file.

Example:
go run main.go  
This runs the program directly in the terminal without generating any output file.
