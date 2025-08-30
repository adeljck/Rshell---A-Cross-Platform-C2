APP_NAME := RShell

BUILD_DIR := build

UPX_OPTS := --best --lzma

PLATFORMS := \
	linux/amd64 \
	windows/amd64 \

LDFLAGS := -s -w

all: clean build compress package finish

build:
	@for platform in $(PLATFORMS); do \
		GOOS=$$(echo $$platform | cut -d/ -f1); \
		GOARCH=$$(echo $$platform | cut -d/ -f2); \
		EXT=$$( [ "$$GOOS" = "windows" ] && echo ".exe" || echo "" ); \
		OUTPUT=$(APP_NAME)-$$GOOS-$$GOARCH$$EXT; \
		echo "🚀 Building $$OUTPUT..."; \
		GOOS=$$GOOS GOARCH=$$GOARCH CGO_ENABLED=1 go build -trimpath -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$$OUTPUT main.go || exit 1; \
	done

compress:
	@for file in $(BUILD_DIR)/*; do \
		if echo $$file | grep -q "darwin"; then \
		  	echo "⚠️ Upx do not support Darwin $$file..."; \
			continue; \
		fi; \
		echo "📦 Compressing $$file..."; \
		upx $(UPX_OPTS) $$file > /dev/null 2>&1 || echo "⚠️ UPX failed for $$file"; \
	done

finish:
	@echo "✨ Build Success To $(BUILD_DIR)"

clean:
	@echo "🧹 Cleaning..."
	@rm -rf $(BUILD_DIR)

.PHONY: all build compress clean finish