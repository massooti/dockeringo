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

