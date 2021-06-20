BIN = bin
OUT = docs
SRC = *.go internal/**/*.go

.PHONY: gen
gen: $(BIN)
	./$(BIN) $(OUT)

.PHONY: local
local: $(BIN)
	./$(BIN) $(OUT) true

$(BIN): $(SRC)
	go build -o $(BIN)

clean:; rm -rf $(BIN) $(OUT)
