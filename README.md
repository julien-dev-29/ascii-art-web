# ascii-art-web

A web-based ASCII art generator written in Go. Users can input text, choose a banner style (standard, shadow, or thinkertoy), and generate ASCII art via a web interface.

## Authors

- Julien

## Usage

### Prerequisites

- Go 1.25 or later

### Running the server

```bash
go run .
```

Or build and run:

```bash
go build -o ascii-art-web
./ascii-art-web
```

The server starts at http://localhost:8000.

### Via Docker

```bash
docker build -t ascii-art-web .
docker run -p 8000:8000 ascii-art-web
```

### Using the web interface

1. Open http://localhost:8000 in your browser
2. Enter text in the textarea
3. Select a banner (Standard, Shadow, or Thinkertoy)
4. Click "Generate"
5. The ASCII art result appears below the form

## Implementation details

### Architecture

The application follows a simple MVC-like structure:

- **main.go** — HTTP server with route handlers (`GET /` and `POST /ascii-art`)
- **generateascii.go** — ASCII art generation logic (model)
- **templates/index.html** — HTML template (view)
- **static/style.css** — CSS styling
- **banners/** — ASCII character set files for each banner style

### Algorithm

1. The server receives text and the selected banner name via a POST form
2. `GenerateAscii` reads the corresponding banner file (`banners/<banner>.txt`)
3. Each banner file contains 95 printable ASCII characters (codes 32-126), each represented as an 8-line glyph separated by a blank line
4. For each character in the input text, the function calculates the starting line index: `(charCode - 32) * 9`
5. It reads 8 lines of the glyph from the calculated offset
6. Lines are concatenated horizontally and separated by newlines for multi-line input
7. The result is passed to the HTML template and rendered inside a `<pre>` block

### HTTP Endpoints

- **GET /** — Returns the main page with the input form
- **POST /ascii-art** — Accepts `text` and `banner` form values, returns the page with generated ASCII art

### HTTP Status Codes

- **200 OK** — Successful request
- **400 Bad Request** — Invalid method, missing text, or invalid banner
- **404 Not Found** — Unknown route
- **500 Internal Server Error** — Banner file not found or template execution failure
