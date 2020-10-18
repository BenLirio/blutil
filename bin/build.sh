cd ./cli-gen
go build
./cli-gen
cd ..
# Manually build utils
go build -v ./cmd/util
