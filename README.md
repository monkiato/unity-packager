# Unity Packager

Unity Packager is a command-line tool providing Unity packages generation without using Unity3D IDE.

It's a very useful tool for CI pipelines, specially for C# libraries that are not strictly designed for Unity or don't required Unity IDE for their development, and can be reused in other C# projects.

## Using the Docker image

Docker images are available under project packages section https://github.com/monkiato/unity-packager/pkgs/container/unity-packager

```
# From your project folder
docker run -v $PWD/:/home/src ghcr.io/monkiato/unity-packager:1.0.0 unity-packager create -p /home/src/Assets -o my-project -i ".csproj"
```

## How to use it

Run `bin/unity-packager` for the help menu

There are different options to customize the way we want to structure our .unitypackage file:

`unity-packager -p <folder-to-package> -o <output-file>`

For instanace:

`unity-package -p Assets/ -o mypackage`

To add ignore patterns:

`unity-package -p Assets/ -o mypackage -i ".csproj"`

To include the `Assets/` folder in case our project doesn't contain it (it's required to be part of the .unitypackage metadata to uncompress the package correctly):

`unity-package -p MyProject/ -o mypackage --add-assets-folder`

## How to generate the binary

Binaries are already available under `bin/` folder, anyway there's a build available to rebuild the code:


`./build.bash`

