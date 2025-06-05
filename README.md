# ASCII-Art

## Overview
ASCII-Art is a Go program that converts an input string into a graphical representation using ASCII characters. The program takes a string as a command-line argument and outputs it in a stylized format based on predefined banner files. It supports letters, numbers, spaces, special characters, and newline (`\n`) characters, rendering each character as an 8-line tall ASCII art.

The project adheres to Go best practices, uses only standard Go packages, and includes functionality to validate input and banner files.

## Objectives
- Convert an input string into an ASCII art representation.
- Handle input containing letters, numbers, spaces, special characters, and newlines (`\n`).
- Use predefined banner files for character templates.
- Validate input to ensure it contains only ASCII characters (0-127).
- Validate banner files using SHA-256 hash checks.
- Ensure the program respects Go coding standards and includes unit tests.

## Installation
1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd ascii-art
   ```

2. **Ensure Go is Installed**:
   - The program requires Go (version 1.16 or later recommended).
   - Install Go from [golang.org](https://golang.org/dl/) if not already installed.

3. **Project Structure**:
   - Place the banner files in the `banners/` directory.
   - The program currently uses `standard.txt` by default.

4. **Dependencies**:
   - Only standard Go packages are used, so no external dependencies are required.

## Usage
Run the program from the command line, providing a single string argument. The output is piped through `cat -e` to display line endings (`$`).

### Command Format
```bash
go run ./cmd "<input-string>" | cat -e
```

### Examples
1. **Empty String**:
   ```bash
   go run ./cmd "" | cat -e
   ```

2. **Newline**:
   ```bash
   go run ./cmd "\n" | cat -e
   $
   ```

3. **Simple Word**:
   ```bash
   go run ./cmd "hello" | cat -e
     _              _   _          $
    | |            | | | |         $
    | |__     ___  | | | |   ___   $
    |  _ \   / _ \ | | | |  / _ \  $
    | | | | |  __/ | | | | | (_) | $
    |_| |_|  \___| |_| |_|  \___/  $
                                   $
                                   $
   ```

4. **Multiple Lines**:
   ```bash
   go run ./cmd "Hello\n\nThere" | cat -e
     _    _          _   _          $
    | |  | |        | | | |         $
    | |__| |   ___  | | | |   ___   $
    |  __  |  / _ \ | | | |  / _ \  $
    | |  | | |  __/ | | | | | (_) | $
    |_|  |_|  \___| |_| |_|  \___/  $
                                    $
                                    $
    $
     _______   _                           $
    |__   __| | |                          $
       | |    | |__     ___   _ __    ___  $
       | |    |  _ \   / _ \ | '__|  / _ \ $
       | |    | | | | |  __/ | |    |  __/ $
       |_|    |_| |_|  \___| |_|     \___| $
                                           $
                                           $
   ```


## File Structure
```
ascii-art/
├── cmd/
│   └── main.go          # Main entry point for the program
├── banners/
│   ├── banners.go       # Banner file parsing and validation logic
│   ├── standard.txt     # Default ASCII art template
    └── checkvalid.go    # Check banner on changing
├── proccesor/
│   └── handling.go    # Logic for processing input and generating ASCII art
└── README.md            # Project documentation
```

## Implementation Details
- **Packages**:
  - `banners`: Handles parsing of banner files and input validation.
    - `Parser`: Reads and parses banner files into a map of runes to 8-line ASCII representations.
    - `IsValidFile`: Validates banner files using SHA-256 hash.
    - `CheckASCII`: Ensures input contains only ASCII characters.
  - `proccesor`: Processes input strings and generates ASCII art output.
  - `main`: Entry point, validates arguments, and orchestrates the program flow.

- **Banner Format**:
  - Each character is represented by 8 lines of ASCII art.
  - Characters are separated by newlines in the banner file.
  - The program currently uses `standard.txt` for rendering.

- **Input Handling**:
  - Supports letters (case-sensitive), numbers, spaces, special characters, and newlines.
  - Non-ASCII characters trigger an error.
  - Newlines (`\n`) in input create separate blocks of ASCII art with appropriate spacing.

- **Error Handling**:
  - Validates the number of command-line arguments.
  - Checks for valid ASCII input.
  - Verifies banner file integrity using SHA-256.

## Testing
To test the code, you can run a prepared script to test the code:

To run tests:
```bash
sh test.sh
```

## Constraints
- Only standard Go packages are used (`fmt`, `log`, `os`, `strings`, `crypto/sha256`, `encoding/hex`).
- Input must be ASCII characters (0-127).
- Banner files must match the expected format and hash.

## Notes
- The program currently uses `standard.txt` as the default banner. To support `shadow.txt` or `thinkertoy.txt`, modify the `Parser` call in `proccesor`. Processing to accept a banner file argument.
- The project is designed to be extensible for additional banner files or features.
- Ensure banner files are present in the `banners/` directory before running the program.

## Example Outputs
The provided code handles various inputs as shown in the usage examples. The output matches the expected format, with each character rendered as an 8-line ASCII art block, separated appropriately for newlines and empty lines.

For further details or to contribute, please refer to the repository or contact the maintainers.
