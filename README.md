## JavaScript Extractor from Source Maps

This Go tool parses JavaScript source map (`.map`) files, extracts the original JavaScript code from the `sourcesContent` field, and saves it as a `.js` file. The tool is designed for scenarios where you need to recover the original, non-minified JavaScript code from a source map.

### Features:
- Parse source map files and extract embedded JavaScript.
- Save the extracted JavaScript to a file or print it to the console.
- Supports handling large source map files efficiently.

### Usage:
1. Build the tool using Go:
   ```bash
   go build -o MapJs MapJs.go
   ```

2. Run the tool:
   - To extract and print the JavaScript code to the console:
     ```bash
     ./MapJs input.map
     ```

   - To extract and save the JavaScript code to a new file:
     ```bash
     ./MapJs input.map output.js
     ```

### Example:
If you have a source map file like `main.js.map`, you can extract the original JavaScript and save it to `main.js` by running:
```bash
./MapJs main.js.map main.js
```

### Requirements:
- Go 1.15 or higher.

### License:
This project is licensed under the GNU General Public License v3.0 (GPLv3). You are free to use, modify, and distribute the software, provided that any modifications you make are also licensed under GPLv3.

For more details, please refer to the full [GPLv3 License](https://www.gnu.org/licenses/gpl-3.0.en.html).

