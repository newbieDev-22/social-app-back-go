# Simple Social App

A simple social app built with Golang, Gin, and Gorm. This app allows users to create accounts, post messages, and comment on posts. It also includes authentication and authorization using JSON Web Tokens (JWT).

## Features

- User authentication and authorization using JWT
- User registration and login
- Post messages
- Comment on posts
- Get a user's profile information

## Technologies Used

- Golang
- Gin
- Gorm
- MySQL
- JWT

## Installation

1. Clone the repository: `git clone https://github.com/natetyler/simple-social-app.git`
2. Install dependencies: `go get -u ./...`
3. Create a MySQL database and update the `DATABASE_URL` environment variable in the `.env` file
4. Run the application: `go run main.go`
5. Open a web browser and navigate to `http://localhost:8888` to access the app

## Endpoints

### User Endpoints

- `POST /register`: Register a new user
- `POST /login`: Login an existing user
- `GET /get-me`: Get the current user's profile information

### Post Endpoints

- `POST /posts`: Create a new post
- `GET /posts/:postId`: Get a post by ID
- `PATCH /posts/:postId`: Update a post
- `DELETE /posts/:postId`: Delete a post
- `GET /posts/:postId/comments`: Get all comments for a post
- `POST /posts/:postId/comments`: Create a new comment for a post
- `GET /posts/:postId/comments/:commentId`: Get a comment by ID
- `PATCH /posts/:postId/comments/:commentId`: Update a comment
- `DELETE /posts/:postId/comments/:commentId`: Delete a comment
