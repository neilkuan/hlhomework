# HL Homework

### Build Go application
```bash
go build -o app
```

### Build image
```
docker build -t go-app .
```


### Use exist image from github container registry
```bash
docker pull ghcr.io/neilkuan/hlhomework:amd64

docker run -d -p 8080:8080 ghcr.io/neilkuan/hlhomework:amd64
```
