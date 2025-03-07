tool = watson-starship

all: buildlocal

buildlocal:
	CGO_LDFLAGS='-static' go build -tags osusergo,netgo -ldflags "-extldflags=-static -w" --trimpath -buildmode=pie -o $(tool)
	strip --strip-all $(tool)

install: buildlocal
	install -m 755 $(tool) $(HOME)/bin/

clean:
	rm -rf $(tool)
