COMMIT = $$(git describe --always)

deps:
	@echo "====> Install dependencies..."
	go get -d github.com/dgraph-io/badger

clean:
	@echo "====> Remove installed binary"
	@echo "====> Remove previous wallets"
	rm -rf tmp
	rm -f chain

install: deps
	@echo "====> Build chain in . "
	mkdir tmp
	go build -ldflags "-X main.GitCommit=\"$(COMMIT)\"" -o chain
