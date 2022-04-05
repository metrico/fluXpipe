<img src="https://user-images.githubusercontent.com/1423657/161866679-e02a438e-2e07-48e4-baf3-6fdbb5187039.png" width=150 />

# FluxPipe
Experimental flux pipeline for embedded or external datasources

### Requirements
- flux library

### Build
``` 
go build -o fluxpipe -ldflags="-s -w" fluxpipe.go
```

### Test
```
cat script.flux | fluxpipe
```
