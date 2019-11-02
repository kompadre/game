PLATFORM := nix
DIST := dist
# For prod, discomment the following line
# GOFLAGS := -ldflags=-w

export GOARCH=amd64
ifeq ($(PLATFORM), win)
	export GOOS=windows
	EXTENSION := -$(PLATFORM)-$(GOARCH).exe 
else
	export GOOS=linux
	EXTENSION := -linux.$(GOARCH)
endif

all: $(DIST)/server$(EXTENSION) $(DIST)/client$(EXTENSION)

$(DIST)/server$(EXTENSION):
	go build $(GOFLAGS) -o $@ cmd/server.go 

$(DIST)/client$(EXTENSION):
	go build $(GOFLAGS) -o $@ cmd/client.go 

clean:
	rm dist/*
