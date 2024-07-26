# TODO List for Notification Microservice

## Everyone can edit and add todo item in this list

## High Priority
1. **[ ] Client connection with WS**
   - **Description**: When Client send handshake(HTTP Upgrade), maybe it's better to use JWT token for that, in that case, we need to validate received JWT token and if it's valid then open connection between Client and Server. Here is basic illustration:
     <img width="971" alt="Screenshot 2024-07-12 at 16 06 32" src="https://github.com/user-attachments/assets/34e19f7c-2c53-4947-8564-a8b6e8a76f13">

## Medium Priority
1. **[ ] Write test cases for WS handler**
   - **Description**: Any testing way is acceptable, you're applied to write test cases using `testify` package, or just use golang build-in package.
    
