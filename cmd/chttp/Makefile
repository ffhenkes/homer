include hserver.env
export $(shell sed 's/=.*//' hserver.env)

all: build test

build:
	$(PACKAGE)
