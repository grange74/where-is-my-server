# Where is my server?

A REST API service in GO to connect up Users and their designated Server.

A simple project to allow me to learn GO. While there are some nice frameworks and libraries out there, my aim is to use the GO built-in libraries wherever possible.

## Tools used for this project

- Using [Sublime Text 2](http://www.sublimetext.com/2) with [GoSublime](https://github.com/DisposaBoy/GoSublime) package for editing GO files.
- Also using a cool utility called [Gin](https://github.com/codegangsta/gin) to reloaded my Web Server code as soon as I modified it. Run using `gin -a "8080" run`.

## Running the project inside a Docker container

I started using the [golang:onbuild](https://registry.hub.docker.com/u/library/golang/) image (~519MB) but switched to using [CenturyLink's golang-builder](https://registry.hub.docker.com/u/centurylink/golang-builder/) as that could create a tiny image (~4.2MB).

Run to build the image:
`docker run --rm -v $(pwd):/src -v /var/run/docker.sock:/var/run/docker.sock centurylink/golang-builder grange74/where-is-my-server`

Run to create a container:
`docker run -d -p 8080:8080 --name="where-is-my-server" grange74/where-is-my-server`