# Railway Seat Allocation

This project provides a web application to allocate train seats based on ticket numbers.

## Seat Arrangement

The seats cycle every 8 tickets:
1. lower
2. middle
3. upper
4. lower
5. middle
6. upper
7. side lower
8. side upper

Then repeats for the next set.

## Setup

### Backend (Golang)

1. Navigate to `backend/` directory.
2. Run `go mod tidy` to download dependencies.
3. Run `go run main.go` to start the server on port 8080.

### Frontend (React)

1. Navigate to `frontend/` directory.
2. Run `npm install` to install dependencies.
3. Run `npm start` to start the development server on port 3000.

## Usage

- Start the backend server.
- Open the frontend in a browser.
- Enter a ticket number and click "Get Seat Position" to see the allocated seat.

## Hosting on GitLab

To host on GitLab:

- For frontend: Build the React app and deploy to GitLab Pages.
- For backend: Use GitLab CI/CD to build and deploy the Golang app, perhaps as a Docker container.

Add a `.gitlab-ci.yml` file for CI/CD pipelines.