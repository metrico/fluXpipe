# Flux as a Library Example


### Build
``` 
go build -o fluxpipe -ldflags="-s -w" fluxpipe.go
```

### Test
```
cat script.flux | fluxpipe
```
