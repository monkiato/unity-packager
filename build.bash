
#!/usr/bin/env bash

package="unity-packager"
    
platforms=("windows/amd64" "windows/386" "darwin/amd64" "darwin/arm64" "linux/386" "linux/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi    

    echo "Building for platform ${platform}"
    env GOOS=$GOOS GOARCH=$GOARCH go build -o bin/$output_name $package
    echo "Build for platform ${platform} finished successfully"
    if [ $? -ne 0 ]; then
           echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done