# gamingwebsite_testtask
Backend developer coding task.

[![Build Status](https://travis-ci.com/Ragnar-BY/gamingwebsite_testtask.svg?branch=master)](https://travis-ci.com/Ragnar-BY/gamingwebsite_testtask)
[![codecov](https://codecov.io/gh/Ragnar-BY/gamingwebsite_testtask/branch/master/graph/badge.svg)](https://codecov.io/gh/Ragnar-BY/gamingwebsite_testtask)

To build with docker use:
```
docker-compose up --build
```


Binary built with 
```
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
```