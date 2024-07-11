# TODO List for Notification Microservice

## High Priority
1. **[ ] Client connection with WS**
   - **Description**: Client connection with WS, it's better to connect client with JWT token, but yet i don't know how is it possible, because it's a security risk..
2. **[ ] Split NATS event functions**
   - **Description**: We need to find place for micro.AddService, now it lives in main.go under the `microServices` function.
    