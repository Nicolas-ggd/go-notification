# Notification Microservice
This microservice is designed to handle system notifications using Go, NATS, and WebSocket. Clients can connect via WebSocket and receive notifications either for all users or for specific clients.

go-notification is a microservice designed to handle system notifications, making sure your users receive important alerts efficiently. With features like WebSocket connectivity, priority-based message handling, and easy integration with NATS, this service is perfect for apps that need robust notification handling.

## Features
- WebSocket Connectivity: It lets clients access the service over WebSocket.
- Broadcast Notifications: Messages are sent to all connected clients.
- Targeted Notifications: Messages are sent to specific clients.
- Priority Notifications: Notifications have different types and priorities (error, warning, info).
- Easy Configuration: Command-line arguments can be used to configure the application.

## Installation
1. Clone repository:

   ```bash
     git clone https://github.com/Nicolas-ggd/go-notification.git
   ```

2. Navigate to the project directory:
   ```bash
     cd go-notifications
   ```
   
3. Install dependencies:

   ```bash
     go mod download
   ```
   
4. Install NATS (if not already installed):

   NATS is a lightweight, high-performance messaging system. You can install NATS using the following methods:

   - Homebrew (macOS):
   ```bash
     brew install nats-server
   ```
   - Docker run:

   ```bash
    docker run -d --name nats-server -p 4222:4222 nats
   ```
   
   Or see full [NATS installation documentation](https://nats.io/download/)

## Configuration
Configure the service using command-line arguments:
- nats-url: The URL of the NATS server (default: nats://nats:4222).
- http-server-port: The port for the HTTP server (default: 8741).

Example:
```bash
  go run ./cmd/gonotification http-server-port=5432 nats-url=nats://127.0.0.1:4222
```

## Usage
Start the service with the following command:
```bash
  go run ./cmd/gonotification http-server-port=5432 nats-url=nats://127.0.0.1:4222
```

Before starting, you can run all tests:
```bash
  make test
```
## Environment Variables
This application uses `.env` file for Github Actions, so you need to configure environment variables, follow these steps:

1. Copy the .env.example file to .env and fill in your environment variables:

   ```shell
   cp .env.example .env
   nano .env
   ```

Check `.env.example` file to known which environment variable you need to deploy applicaiton

2. After creating environment file, load the environment variables:

   ```shell
   ./load_env.sh
   ```
   
## Sending Notifications

1. Command line notification request using NATS:
   Send a message to all connected clients

   ```shell
   nats req NOTIFICATION.send-to-all '{"type": "error", "message": "example", "time": "2024-04-17T09:00:00Z"}'
   ```

   Send a message to specific clients

   ```shell
   nats req NOTIFICATION.send-to-clients '{"type": "warning", "message": "example", "time": "2024-04-17T09:00:00Z", "clients": ["1", "2"]}'
   ```
   
2. Broadcast Notification:
   Send a message to all connected clients.

   ```json
   {
     "type": "error",
     "message": "This is a broadcast go-notification.",
     "time": "2024-04-17T09:00:00Z"
   }
   ```
   
   #### Response JSON should be like:
   ```json
    {
      "id": 1,
      "type": "error",
      "message": "This is a broadcast go-notification.",
      "time": "2024-04-17T09:00:00Z"
    }
   ```

3. Targeted Notification:
   Send a message to specific clients.

   ```json
   {
     "type": "warning",
     "clients": ["1", "2"],
     "message": "This is a targeted go-notification.",
     "time": "2024-04-17T09:00:00Z"
   }
   ```

   #### Response JSON should be like:
      ```json
       {
         "id": 1,
         "type": "warning",
         "message": "This is a broadcast go-notification.",
         "time": "2024-04-17T09:00:00Z"
       }
      ```

## Notification Types and Priorities
There are three types of system notifications, each with a different priority level:
- Error: Highest priority. These notifications are sent first.
- Warning: Medium priority. These notifications are sent after error notifications.
- Info: Lowest priority. These notifications are sent last.

If there is a list of notification that are received simultaneously, the client will receive this list in the following order as [described](#notification-types-and-priorities)

This illustration shows Notification list before and after sorting according to priority types:
<img width="971" alt="Screenshot 2024-07-15 at 12 27 46" src="https://github.com/user-attachments/assets/ecc4947b-9b78-4e07-91b5-2273090a8a16">

- Notification whose type is error is evaluated with priority 1
- Notification of warning type is evaluated with priority 2
- Notification of info type is evaluated with priority 3

**Note**: that ordering and sending notifications by priority only works when you send multiple notifications in the application at the same time.

## Example Client
Here's a simple example of a client connecting to the WebSocket server and handling messages:

**NOTE**: Client side isn't stable yet, application is during the development
```javascript
const socket = new WebSocket('ws://localhost:8080/ws');

socket.onopen = function(event) {
  console.log('Connected to WebSocket server.');
};

socket.onmessage = function(event) {
  const message = JSON.parse(event.data);
  console.log('Received:', message);
};

socket.onclose = function(event) {
  console.log('Disconnected from WebSocket server.');
};

socket.onerror = function(error) {
  console.error('WebSocket Error:', error);
};
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License
This project is licensed under the MIT License.
