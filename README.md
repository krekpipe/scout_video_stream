

# JPEG Stream Server for Moorebot Scout

## Overview

The JPEG Stream Server is a Go application specifically designed for the Moorebot Scout robot. It subscribes to a ROS (Robot Operating System) topic to receive JPEG frames from the Moorebot Scout's camera system and serves these frames over an MJPEG HTTP stream. This setup allows you to stream video data from the Moorebot Scout via a web browser or other HTTP clients.

## Features

- **ROS Integration:** Subscribes to ROS topics to receive image frames from Moorebot Scout.
- **MJPEG Streaming:** Serves the received JPEG frames as an MJPEG stream over HTTP.
- **Configurable:** Supports setting ROS and local host addresses through command-line flags.

## Requirements

- **Go:** Version 1.16 or higher (for building from source)
- **ROS:** ROS Master should be accessible
- **GOROSLIB:** Go library for interacting with ROS

## Installation

### Downloading Pre-built Binary

You can download a pre-built binary for Moorebot Scout from the [releases page](https://github.com/yourusername/jpeg-stream-server/releases).

1. **Go to the [Releases Page](https://github.com/yourusername/jpeg-stream-server/releases).**
2. **Download the appropriate binary for your operating system:**
   - For Linux: `jpeg-stream-server-linux-amd64.tar.gz`
   - For macOS: `jpeg-stream-server-darwin-amd64.tar.gz`
   - For Windows: `jpeg-stream-server-windows-amd64.zip`
3. **Extract the downloaded archive:**

   - **Linux/macOS:**

     ```sh
     tar -xzvf jpeg-stream-server-linux-amd64.tar.gz
     ```

   - **Windows:**

     Extract the `.zip` file using a tool like WinRAR or 7-Zip.

4. **Navigate to the extracted directory:**

   ```sh
   cd path/to/extracted/directory
   ```

### Building from Source

If you prefer to build from source, follow these steps:

1. **Clone the Repository**

   ```sh
   git clone https://github.com/yourusername/jpeg-stream-server.git
   cd jpeg-stream-server
   ```

2. **Build the Application**

   ```sh
   go build -o jpeg-stream-server
   ```

## Usage

To run the server, use the following command:

```sh
./jpeg-stream-server -h <ROS_HOST> -l <LOCALHOST>
```

- `-h <ROS_HOST>`: Address of the ROS master in the format `IP_ADDRESS:PORT`. Default is `192.168.1.61:11311`.
- `-l <LOCALHOST>`: Local address to bind the HTTP server. Default is `127.0.0.1`.

### Example

```sh
./jpeg-stream-server -h 192.168.1.61:11311 -l 127.0.0.1
```

This command will start the server with ROS master at `192.168.1.61:11311` and the HTTP server listening on `127.0.0.1:8080`.

## Accessing the Stream

Once the server is running, you can access the MJPEG stream at:

```
http://127.0.0.1:8080/stream
```

## Configuration

The application supports configuration via command-line flags:

- `-h` or `--host`: ROS master address.
- `-l` or `--localhost`: Localhost address for the HTTP server.

## Code Explanation

### Frame Message Type

Defines the structure of the ROS messages containing image data from the Moorebot Scout's camera. It includes fields for metadata and the image data itself.

### Main Function

- **Node Creation:** Initializes a ROS node with the given configuration.
- **Subscriber Creation:** Subscribes to the ROS topic `/CoreNode/jpg` and sets a callback to handle incoming messages.
- **HTTP Server:** Starts an HTTP server to serve MJPEG streams.
- **Signal Handling:** Listens for interrupt signals to allow graceful shutdown.

### Callback Function

Handles incoming ROS messages by sending the image data to a channel.

### HTTP Handler

Handles HTTP requests to serve MJPEG streams. It reads image data from the channel and writes it to the HTTP response in the MJPEG format.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests for any enhancements or bug fixes.

## Contact

For any questions or issues, please open an issue on the GitHub repository or contact [your.email@example.com](mailto:your.email@example.com).

---

Feel free to adjust URLs, file names, or any other specific details to match your project and distribution setup.
