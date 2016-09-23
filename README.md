# fabric8menu

## setup

Using glide hasn't pwrked so far, for now clone this repos and:
```
go get github.com/alexflint/gallium
```


## Builing 

Make file not working yet
```
go run fabric8menu.go
```

TODO sort out Makefile and figure out how to prevent menu and not status in nav bar
//go build -o build/fabric8menu  ./fabric8menu.go && gallium-bundle -o fabric8menu.app build/fabric8menu  && open fabric8menu.app 