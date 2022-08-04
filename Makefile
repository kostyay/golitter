
TINYGO=/Users/kostyay/Downloads/tinygo/bin/tinygo
UNO_DEV_PATH=$(shell ls /dev/cu.usb*)
OUTPUT_PATH=./output.hex
BUILD_TARGET=arduino
BOARD=arduino:avr:uno

.PHONY: deploy

build:
	@echo "Building image to $(OUTPUT_PATH)"
	$(TINYGO) build -scheduler=tasks -target=$(BUILD_TARGET) -o $(OUTPUT_PATH) ./main.go

deploy:
	@echo "Deploying image to $(UNO_DEV_PATH)"
	@arduino-cli upload -b $(BOARD) -p $(UNO_DEV_PATH) -i $(OUTPUT_PATH) .

all: build deploy
	@echo "Done"
