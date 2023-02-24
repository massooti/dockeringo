# Dockeringo

Dockeringo is a Golang application that allows you to interact with Docker containers interactively. With Dockeringo, you can start, stop, and manage Docker containers easily through a command-line interface.

Requirements

* Docker
* Golang


# Installation


Clone the repository:
```
git clone https://github.com/massooti/dockeringo
```

Navigate to the root directory of the Dockeringo repository:

```
cd dockeringo
```

Build the Docker image using the Dockerfile:

```
docker build -t dockeringo .
```

Run the Docker container:

```
docker run -it --rm dockeringo
```

