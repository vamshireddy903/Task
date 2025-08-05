# Go Web App CI/CD with Jenkins and Docker
This project demonstrates a simple Go (Golang) web application deployed using a Jenkins CI/CD pipeline with Docker.

# Project Overview
A basic Go web server that serves a styled HTML page.

CI/CD setup using Jenkins to automate:

Clone the GitHub repo

Build the Go app

Run tests 

Build a Docker image

Deploy the Docker container exposing port 5000

 #Files Included
main.go: Go source code that serves a simple HTML page.

Dockerfile: Multi-stage Docker build for efficient and small image size.

Jenkinsfile: Pipeline script for automating CI/CD stages (clone, build, test, deploy).

Port 5000 is exposed and used for the Go web server.

# How It Works
Jenkins pulls code from GitHub on push via webhook:

Jenkins executes the pipeline:

Build the Go app

Dockerize the app

Run the Docker container on port 5000

# Accessing the App
Once deployed, you can access the app via:

http://<EC2_PUBLIC_IP>:5000

You'll see:
<img width="1891" height="955" alt="image" src="https://github.com/user-attachments/assets/15130e00-7e7b-4ea4-8169-13ea52ca45f5" />

<img width="1889" height="664" alt="image" src="https://github.com/user-attachments/assets/892d48af-30ea-4424-bf88-4359a02a4af5" />


<img width="1861" height="956" alt="image" src="https://github.com/user-attachments/assets/6a295adb-600f-40d3-be77-e5ac2a87c37d" />

<img width="1919" height="893" alt="image" src="https://github.com/user-attachments/assets/97eb6bf5-1dee-46dc-b9d9-c151c43a20f0" />


# Technologies Used

Docker

Jenkins

GitHub 

EC2 (for hosting Jenkins and Docker)



