# go-cutup..

this little app implement the cut-up technique, described [here](https://en.wikipedia.org/wiki/Cut-up_technique).

## docker

build with the Dockerfile.scratch

```
docker build -t go-cutup:0.0.2 . -f Dockerfile.scratch
```

run the image

```
docker run -p 3334:3334 go-cutup:0.0.2
```
