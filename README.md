# Tunnel Daemon

Tunnel Daemon is a Go application that manages Linux GRE tunnels using the `ip` command. It provides a RESTful API to create, delete and list tunnels.

## Requirements

- Go 1.16+
- Linux

## Installation

1. Clone the repository:

`git clone https://github.com/Digital-Servers/tunnel-daemon.git
cd tunnel-daemon`

2. Build the application:

`go build`

3. Run the application:

`./tunnel-daemon`

## Usage

The application exposes a RESTful API that allows the user to create, delete and list tunnels.

### Creating a Tunnel

To create a new tunnel, send a `POST` request to the `/tunnel` endpoint with the following parameters:

- `localIP`: the local IP address of the tunnel
- `tunnelName`: the name of the tunnel
- `remoteIP`: the remote IP address of the tunnel

Example:

`POST /tunnel HTTP/1.1
Host: localhost:8080
Content-Type: application/x-www-form-urlencoded`

`localIP=10.0.0.1&tunnelName=mytunnel&remoteIP=192.168.0.1`

If the tunnel is created successfully, the server will respond with an HTTP 200 OK status code and a JSON object containing a success message.

If there is an error creating the tunnel, the server will respond with an HTTP 500 Internal Server Error status code and a JSON object containing an error message.

### Deleting a Tunnel

To delete an existing tunnel, send a `DELETE` request to the `/tunnel/:name` endpoint, where `:name` is the name of the tunnel to be deleted.

Example:

`DELETE /tunnel/mytunnel HTTP/1.1
Host: localhost:8080`

If the tunnel is deleted successfully, the server will respond with an HTTP 200 OK status code and a JSON object containing a success message.

If there is an error deleting the tunnel, the server will respond with an HTTP 500 Internal Server Error status code and a JSON object containing an error message.

### Listing Tunnels

To list all existing tunnels, send a `GET` request to the `/tunnels` endpoint.

Example:

`GET /tunnels HTTP/1.1
Host: localhost:8080`

If the tunnels are listed successfully, the server will respond with an HTTP 200 OK status code and a JSON object containing a map of the tunnel names and peer IP addresses.

If there is an error listing the tunnels, the server will respond with an HTTP 500 Internal Server Error status code and a JSON object containing an error message.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
