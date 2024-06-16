# Go Money

Go-Money gives users the ability to manage their incomings and outgoings in an easy and effective way.

![finance-app](https://github.com/user-attachments/assets/0552006e-dc91-4664-a017-00241c1c7bd8)

This app is built using go serving as the back and Next JS and typescript as the front.

# Setting up the ports

I've set server to listen on port to :8080, so change this if needed.

if you do change the port from :8080 dont forget to amend the port in the client api folder.

Allowing localhost:3000 for the client, so change this if needed

# How to configure DB connection

use the example .env

navigate to database.go and add your DB credentials

# How to run the app

```
Back:
go mod tidy
cd /cmd

from root run make run

```

```
Front:
cd /web/client

run npm install -g pnpm
run pnpm install
run pnpm dev

docker:

from root of project where docker-compose.yml is located run the following
docker-compose up --build

```

I decided to use Gorm to connect to the db and query, since 1.22 and the change to the native http package in the way you define routes, I decided to use this instead of chi or fibre. golang-jwt was used for authentication.

On the front, used react hook form, zod, axios, tanstack query/table, tailwind and shadcn.

This is still a work in progress...

<img width="494" alt="Screenshot 2024-12-16 at 11 21 01 AM" src="https://github.com/user-attachments/assets/e4f21fba-ca0e-4793-b933-b629d708c018" />

<img width="437" alt="Screenshot 2024-12-16 at 2 46 39 PM" src="https://github.com/user-attachments/assets/6ca658b5-de2d-458b-91b5-e63f479bc58f" />

<img width="995" alt="Screenshot 2024-12-18 at 6 45 01 AM" src="https://github.com/user-attachments/assets/d96667cc-c519-4a69-8995-d9b5633d6cbc" />







