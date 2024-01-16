default: build

version = 3.9.0-1
install_path = ~/.terraform.d/plugins/terraform.local/local/pcluster/$(version)
file_name = terraform-provider-pcluster_$(version)

.PHONY: testacc darwin_arm64 darwin_amd64 windows_amd64 linux_amd64 all install
darwin_arm64:
	env GOOS=darwin GOARCH=arm64 go build -o build/darwin_arm64
darwin_amd64:
	env GOOS=darwin GOARCH=amd64 go build -o build/darwin_amd64
windows_amd64:
	env GOOS=windows GOARCH=amd64 go build -o build/windows_amd64
linux_amd64:
	env GOOS=linux GOARCH=amd64 go build -o build/linux_amd64

install_darwin_arm64: darwin_arm64
	mkdir -p $(install_path)/darwin_arm64
	install build/darwin_arm64 $(install_path)/darwin_arm64/$(file_name)
install_darwin_amd64: darwin_amd64
	mkdir -p $(install_path)/darwin_amd64
	install build/darwin_amd64 $(install_path)/darwin_amd64/$(file_name)
install_windows_amd64: windows_amd64
	mkdir -p $(install_path)/windows_amd64
	install build/windows_amd64 $(install_path)/windows_amd64/$(file_name)
install_linux_amd64: linux_amd64
	mkdir -p $(install_path)/linux_amd64
	install build/linux_amd64 $(install_path)/linux_amd64/$(file_name)

all: darwin_arm64 darwin_amd64 windows_amd64 linux_amd64

install: install_darwin_arm64 install_darwin_amd64 install_windows_amd64 install_linux_amd64

# Run acceptance tests

test:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

testend2end:
	TF_ACC=1 go test ./... -v -run="^TestEnd2End" $(TESTARGS) -timeout 120m

testunit:
	TF_ACC=1 go test ./... -v -run="^TestUnit" $(TESTARGS) -timeout 10m -cover
