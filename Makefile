SHELL=/bin/bash

.PHONY: $(shell egrep -oh ^[a-zA-Z0-9][a-zA-Z0-9_-]+: $(MAKEFILE_LIST) | sed 's/://')

set-up:
	docker-compose build

up:
	docker-compose up

up-d:
	docker-compose up -d
