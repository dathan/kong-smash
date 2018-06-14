# kong-smash
This repository is to break kong with HTTP requests. There will be a server to sleep and a driver to destroy

```
./smash -h
Usage of ./smash:
 -concurrency int
  	how many concurrent processes to run. Raise this to smash more (default 4)
 -delay int
  	add a delay flag to cause the service to block in ms time
 -requests int
  	how many requests to perform (default 1000)
 -url string
  	url to smash (default "http://localhost:8282/simulate")
```
