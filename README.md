# Boot.dev CICD Course

https://www.boot.dev/courses/learn-ci-cd-github-docker-golang

This was a good exercise to work through the concepts of CI/CD. It uses Google Cloud Run for the application and Turso for the database.

There was nothing new for me in this course. But it was really good to stitch everything together and deploy a real application with database migrations.

I did not know about Google Cloud Run, that looks like a good option for deploying applications. Although the admin UI is slow and not very user friendly -- but still way better than AWS.

What I liked:

- The course is well structured and has good pace.
- Google Cloud Run is a good option for deploying applications.
- It takes you through all the necessary steps to configure Google Cloud services. In real life this admin work is the most time consuming part.

What I did not like:

- The instructions where lacking in some places (eg. tagging Docker images, navigating Google Cloud UI). But this could also be seen as a good thing because it forces you to figure it out yourself.
- The course stopped a bit short. When main auto deploys on push, branch protection rules should be added to prevent accidental pushes.
- Environments like test, prod, staging are not mentioned. A side effect of using environments is that the status shows up in the sidebar of the repo home page.

Overall: Would recommend!

My referral link: https://www.boot.dev?bannerlord=hanspetter

---

## learn-cicd-starter (Notely)

![CI badge](https://github.com/HP/go-cicd/actions/workflows/ci.yml/badge.svg)

This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

## Local Development

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

Run the server:

```bash
go build -o notely && ./notely
```

_This starts the server in non-database mode._ It will serve a simple webpage at `http://localhost:8080`.

You do _not_ need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!

HP's version of Boot.dev's Notely app
