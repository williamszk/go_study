# Objective

Implement a simple static page with a Go back-end.

We are using the code from another project: `220702_05_go_projects`

The objective here is to reproduce the application but running it inside containers.

In this project we are not using any web back end framwork in Go
we are using just the simple package net/http

Given that we don't have many containers, we can just use a simple Dockerfile to build our image


Tasks:
- Containarize application OK
- Push image to AWS ECR
