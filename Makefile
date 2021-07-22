BIN = bin
OUT = docs
POST = $(OUT)/posts
SRC = *.go internal/**/*.go

.PHONY: gen
gen: $(BIN)
	./$(BIN) $(OUT) $(POST)

.PHONY: local
local: $(BIN)
	./$(BIN) $(OUT) $(POST) true

$(BIN): $(SRC)
	go build -o $(BIN)

clean:; rm -rf $(BIN) $(OUT)
