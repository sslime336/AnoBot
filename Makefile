.DEFAULT_GOAL:=bot

BOOT_FILE:=main.go
OUTPUT:=./AnoBot.exe

DEVICE_INFO:=./device.json
SESSION_TOKEN:=./session.token
QRCODE_PNG:=./qrcode.png

.PHONY: bot
bot: $(BOOT_FILE) clean
	@go build -o $(OUTPUT) $(BOOT_FILE)

.PHONY: release
release: clean
	@go build -o $(OUTPUT) -ldflags='-s -w' $(BOOT_FILE)

.PHONY: run
run: bot
	@$(OUTPUT)

.PHONY: debug
debug: bot
	@dlv exec $(OUTPUT)

.PHONY: clean
clean:
	@if [ -e $(OUTPUT) ]; then rm $(OUTPUT); fi
	@if [ -e $(QRCODE_PNG) ]; then rm $(QRCODE_PNG); fi

.PHONY: clean_all
clean_all: clean
	@if [ -e $(DEVICE_INFO) ]; then rm $(DEVICE_INFO); fi
	@if [ -e $(SESSION_TOKEN) ]; then rm $(SESSION_TOKEN); fi
