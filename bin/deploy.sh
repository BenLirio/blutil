# Manually build utils
go build -v ./blutil
gsutil rm gs://lirio-tools-release/blutil
gsutil cp ./blutil gs://lirio-tools-release/blutil
rm ./blutil
