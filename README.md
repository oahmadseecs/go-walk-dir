# GO WALK DIR

A program that traverses the directory and it's subdirectories. It collects the permissions of each file and directory inside the root directory provided, and creates a set of permissions observed during the traversal.

## Usage

### Requirements

```go version go1.18.3```

### Build the Binary

```bash
go build -o walk-dir(.exe)
```

Build the binary using this command. `.exe` in the parenthesis is only needed on windows platform. This command will create a binary in current working directory.

### Use the Binary

```bash
walk-dir(.exe) [DIRECTORY_LOCATION] [OUTPUT_FILENAME]
```

## Results

Results are stored in a file created in the current working directory.

### File Template

```text
Maximum Size of a File in Directory: [SIZE] MB
Minimum Size of a File in Directory: [SIZE] MB
All Directory Permissions encountered: 
d--------- [LIST] where - : r,w,x
All Files Permissions encountered: 
---------- [LIST] where - : r,w,x
```
