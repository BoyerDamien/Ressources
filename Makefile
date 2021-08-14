DOC_FILENAME = specs.json
DOC_GEN = swagger generate spec -w $(PWD) -o $(DOC_FILENAME) -m

MARDOWN_FILENAME = readme.md
MARKDOWN_GEN = swagger generate markdown -f $(DOC_FILENAME) --output $(MARDOWN_FILENAME)
RM = rm -rf

doc: install
	$(DOC_GEN)
	$(MARKDOWN_GEN)

run: build
	swagger serve $(DOC_FILENAME)

install:
	which swagger || go get -u github.com/go-swagger/go-swagger
	go mod tidy

rerun: re run

re: clean doc


clean:
	go clean
	$(RM) $(DOC_FILENAME) $(MARDOWN_FILENAME)
	find . -name "test.db" -exec rm {} \;

test:
	go test -v ./...
