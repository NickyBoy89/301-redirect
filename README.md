# Info

This is an extremely compact Docker image that redirects users with a 301 redirect

It is written in Golang, and uses multi-stage builds to build and than copy the binary into a scratch container of around 6.4MB

# Building

1. Clone the repo and `cd` into the folder

1. `docker build .`