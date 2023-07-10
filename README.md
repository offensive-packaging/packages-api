# packages-api
## Run it locally

You can try this program by running it in its interpreter mode via 

````
$ go run main.go
````

or build it with 

````
$ go build -o packages-api
````

Then just make a GET request to the following address, as in the example: 

````
$ http://localhost:8080/packages?branch=main&arch=amd64&package=nginx
````

or just use cURL or Postman.

## Options 

The options for branch and arch (architecture) are:

- branch: `main, contrib, non-free`
- arch: `amd64, arm64, armhf, i386`
