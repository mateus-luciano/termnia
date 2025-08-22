APP := termnia
BIN := bin


.PHONY: tidy run


tidy:
go mod tidy


run:
go run .


# === Cross‑compile quick builds (require Go toolchain; Fyne can run without cgo on most setups) ===


win:
GOOS=windows GOARCH=amd64 go build -o $(BIN)/$(APP)-windows-amd64.exe .


linux:
GOOS=linux GOARCH=amd64 go build -o $(BIN)/$(APP)-linux-amd64 .


mac:
# Choose one arch (arm64 for Apple Silicon, amd64 for Intel)
GOOS=darwin GOARCH=arm64 go build -o $(BIN)/$(APP)-darwin-arm64 .


# === Packaging with Fyne CLI (creates .exe/.app/.desktop) ===
# First: go install fyne.io/fyne/v2/cmd/fyne@latest


package-win:
fyne package -os windows -icon icon.png -name "Termnia"


package-mac:
fyne package -os darwin -icon icon.png -name "Termnia"


package-linux:
fyne package -os linux -icon icon.png -name "Termnia"