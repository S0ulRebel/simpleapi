You can switch database type in the .env file (postgres or mongodb)

Create user is basically registration and doesn't need login

When you create a user and login you can fetch update and delete users (no complex verification or roles implemented as this is a test project)
You can then create posts.

All operations except create user and login require you to be logged in.

Wrote only a few unit tests for one service as it's not a real world project.

Postman collection provided in the project root.
