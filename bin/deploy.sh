# Manually build utils
go build -v ./cmd/util
gsutil rm gs://lirio-tools-release/util
gsutil cp ./util gs://lirio-tools-release/util
rm ./util
