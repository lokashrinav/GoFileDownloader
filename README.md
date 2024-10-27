# Concurrent File Downloader

A simple Go application that downloads multiple files concurrently from specified URLs. This application demonstrates the use of Goroutines, WaitGroups, and channels for handling concurrent tasks in Go.

## Features

- **Concurrent Downloads**: Downloads files concurrently, significantly improving download efficiency.
- **HTTP Request Handling**: Uses Go's standard `net/http` library for making HTTP requests.
- **File Handling**: Utilizes Go's `os` and `io` packages for file operations.
- **Concurrency Management**: Employs Goroutines and WaitGroups to manage concurrent tasks.
- **Error Handling**: Includes basic error handling for network requests and file operations.
- **Performance Metrics**: Measures and displays the total time taken for all downloads.
- **Progress Indication**: Optionally displays a progress bar for each download (using a library like `github.com/schollz/progressbar/v3`).

## Requirements

- **Go Version**: Go 1.16 or later
- **Internet Connection**: Required to download files from specified URLs

## Getting Started

### Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/concurrent-file-downloader.git
   cd concurrent-file-downloader
   ```

2. **Build the Application**:
   ```bash
   go build
   ```

### Usage

1. **Update the File URLs**:
   Open the `main.go` file and update the `fileURLs` variable with the URLs of the files you want to download.
   ```go
   var fileURLs = []string{
       "URL_1",
       "URL_2",
       "URL_3",
   }
   ```

2. **Run the Application**:
   ```bash
   ./concurrent-file-downloader
   ```
   The files will be downloaded to the current directory, and the total time taken will be displayed in the console.

## Code Explanation

### Main Function
- Initializes a `WaitGroup` to manage concurrent downloads.
- Records the start time to measure the total download time.
- Iterates over the list of file URLs, starting a new Goroutine for each URL.

### Download Function
- Fetches the file from the provided URL using `http.Get()`.
- Creates a new file and copies the contents using `io.Copy()`.
- Closes the response body to avoid resource leaks.
- Handles errors during network requests and file operations.

### Concurrency and Error Handling
- Uses Goroutines to download files concurrently.
- Employs channels for communication between Goroutines to handle success and error messages.
- Utilizes a `WaitGroup` to ensure the program waits for all Goroutines to finish before exiting.

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.

## Author

Shrinav Loka - [Your GitHub Profile](https://github.com/yourusername)

## Contributions

Feel free to contribute to this project by opening issues or submitting pull requests. Your contributions are highly appreciated and will help improve the functionality and robustness of this concurrent file downloader.

## Additional Considerations

- **Chunking and Resuming**: For large files, consider splitting the download into smaller chunks and resuming downloads from where they left off to handle network failures and improve efficiency.
- **Progress Bar**: Implement a progress bar to show the download progress, enhancing the user experience.
- **Memory Efficiency**: Ensure memory efficiency by directly writing chunks to the file on disk instead of loading them into memory.
- **Timeouts and Connection Limits**: Set timeouts for HTTP requests and limit the number of concurrent connections to avoid issues with poorly implemented servers and to prevent potential DOS attacks.