# Dev Environment
### Container
```
docker build --target build . -t go
docker run --publish 8080:8080 -it -v ${PWD}:/work go sh
```