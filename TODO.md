# TODO List for Notification Microservice

## High Priority
1. **[ ] Client connection with WS**
   - **Description**: When Client send handshake(HTTP Upgrade), maybe it's better to use JWT token for that, in that case, we need to validate received JWT token and if it's valid then open connection between Client and Server. Here is basic illustration:
2. **[ ] Split NATS event functions**
   - **Description**: We need to find place for micro.AddService, now it lives in main.go under the `microServices` function.
    