<img src="https://user-images.githubusercontent.com/1423657/161867564-4b3fc400-95e5-424c-9210-604d5671a85e.png" width=100 />

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
