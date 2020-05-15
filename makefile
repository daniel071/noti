# tool marcros
CC := go build
CCFLAG :=

# path marcros
BIN_PATH := bin
TARGET := noti

$(TARGET): main.go
    $(CC) $(CCFLAG) -o $(BIN_PATH)/$(TARGET) $?
