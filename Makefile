DOC_FILENAME = specs.json
DOC_GEN = swagger generate spec -w $(PWD) -o $(DOC_FILENAME) -m

MARDOWN_FILENAME = readme.md
MARKDOWN_GEN = swagger generate markdown -f $(DOC_FILENAME) --output $(MARDOWN_FILENAME)
RM = rm -rf

build: install doc
	
run: build
	swagger serve $(DOC_FILENAME)

install:
	which swagger || go get -u github.com/go-swagger/go-swagger
	go mod tidy

rerun: re run


clean:
	go clean
	$(RM) $(DOC_FILENAME) $(MARDOWN_FILENAME) test.db

doc: install
	$(DOC_GEN)
	$(MARKDOWN_GEN)

re: clean build

test:
	go test -v ./...
