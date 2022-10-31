NAME="main"
oses=("darwin" "linux")
arches=("amd64" "arm64")
for os in ${oses[@]}; do
  for arch in ${arches[@]}; do
    filename=$NAME-$os-$arch
      GOOS=$os GOARCH=$arch go build -o $filename
  done
done
