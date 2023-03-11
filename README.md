# Dockeringo

This is an application written golang which make docker containers from [Alobotpy](http://gitlab.kavano.org/backend-phyton/alobotpy) dynamically.

## Requirments & Dependancies

-   Golang 1.18



##Installation 

via docker `docker build .` && `docker run -d`


via manually `go run build`


## How it works?

for starting bot use this end point 


endpoint: 

 - `<localhost:9808>/start `

params:
```
link: https://test.alocom.co/class/kavano/daily

guests:100
```

for stoping bot use this end point 


endpoint: 

 - `<localhost:9808>/stop `
 - 
 
 
## NOTICE: After using the application once all of your works are finished you should stop the robot using API.
=======
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

