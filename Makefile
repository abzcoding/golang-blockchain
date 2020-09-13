COMMIT = $$(git describe --always)

deps:
	@echo "====> Install dependencies..."
	go get -d github.com/dgraph-io/badger
	go get -d github.com/mr-tron/base58
	go get -d golang.org/x/crypto/ripemd160

clean:
	@echo "====> Remove installed binary"
	@echo "====> Remove previous wallets"
	rm -rf tmp
	rm -f chain

install: deps
	@echo "====> Build chain in . "
	mkdir -p tmp
	go build -ldflags "-X main.GitCommit=\"$(COMMIT)\"" -o chain
