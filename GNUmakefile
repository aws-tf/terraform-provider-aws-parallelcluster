default: build

ifndef VERSION
VERSION = 1.1.0
endif

install_path = ~/.terraform.d/plugins/terraform.local/local/aws-parallelcluster/$(VERSION)
file_name = terraform-provider-aws-parallelcluster_$(VERSION)

.PHONY: testacc darwin_arm64 darwin_amd64 windows_amd64 linux_amd64 all install

# Build
darwin_arm64:
	env GOOS=darwin GOARCH=arm64 go build -o build/darwin_arm64
darwin_amd64:
	env GOOS=darwin GOARCH=amd64 go build -o build/darwin_amd64
windows_amd64:
	env GOOS=windows GOARCH=amd64 go build -o build/windows_amd64
linux_amd64:
	env GOOS=linux GOARCH=amd64 go build -o build/linux_amd64

build: darwin_arm64 darwin_amd64 windows_amd64 linux_amd64

# Install
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

install: install_darwin_arm64 install_darwin_amd64 install_windows_amd64 install_linux_amd64

# Uninstall
uninstall_darwin_arm64:
	rm -rf $(install_path)/darwin_arm64/$(file_name)
uninstall_darwin_amd64:
	rm -rf $(install_path)/darwin_amd64/$(file_name)
uninstall_windows_amd64:
	rm -rf $(install_path)/windows_amd64/$(file_name)
uninstall_linux_amd64:
	rm -rf $(install_path)/linux_amd64/$(file_name)

uninstall: uninstall_darwin_arm64 uninstall_darwin_amd64 uninstall_windows_amd64 uninstall_linux_amd64

# Clean
clean:
	rm -rf build/*

# Run tests
test:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
	cd internal/provider/openapi; \
	TF_ACC=1 go test ./... -v -run="^Test_openapi" $(TESTARGS) -timeout 10m -cover

test_end2end:
	TF_ACC=1 go test ./... -v -run="^TestEnd2End" $(TESTARGS) -timeout 120m

test_end2end_cluster:
	TF_ACC=1 go test ./... -v -run="^TestEnd2EndCluster" $(TESTARGS) -timeout 120m

test_end2end_image:
	TF_ACC=1 go test ./... -v -run="^TestEnd2EndImage" $(TESTARGS) -timeout 120m

test_unit:
	TF_ACC=1 go test ./... -v -run="^TestUnit" $(TESTARGS) -timeout 10m -cover

test_generated:
	cd internal/provider/openapi; \
	TF_ACC=1 go test ./... -v -run="^Test_openapi" $(TESTARGS) -timeout 10m -cover
