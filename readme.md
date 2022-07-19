# Beeper

A go toast notification package for Windows, Linux and macOS

Credits to:
- [jacobmarshall/pokevision-cli](https://github.com/jacobmarshall/pokevision-cli).
- [go-toast/toast](https://github.com/go-toast/toast)
- [gen2brain/beeep](https://github.com/gen2brain/beeep)

## CLI

I found several projects that partially offer the features that I need. The ability to pass custom messages and image paths as params. 


```cmd
flag needs an argument: -img & -msg
Usage of C:\Example\beeper\main.go-windows-amd64.exe:
  -img string
        Enter a Custom Path (default "assets/information.png")
  -msg string
        Enter a Custom Message (default "Default BEEP BEEP!")
  -title string
        Enter a Custom Title (default "Title")

```

## Example
This assumes you are familar with go and have a god development environment setup. The below was performed on windows, however thsi can be performed on macOS or Linux, change the commands as necessary. 

1. Make Directories and CD to beeper parent dir
```cmd
C:\Example> mkdir beeper && mkdir beeper\assets
C:\Example> cd beeper
```

2. Add Images to assets dir

3. Create main go package
```cmd
C:\Example\beeper> echo package main  > main.go
```

4. Add Example Code
```go
package main

import (
	"flag"
	"fmt"
	"reflect"

	"github.com/gen2brain/beeep"
)

func main() {
	//---Define the various flags---
	msgPtr := flag.String("msg", "Default BEEP BEEP!",
		"Enter a Custom Message")
	imgPth := flag.String("img", "assets/information.png",
		"Enter a Custom Path")
    titleNm := flag.String("title", "Title",
    "Enter a Custom Title")
	//---parse the command line into the defined flags---
	flag.Parse()
	err := beeep.Notify(*titleNm, *msgPtr, *imgPth)
	if err != nil {
		panic(err)
	}
	for _, arg := range flag.Args() {
		fmt.Print(arg, " ")
		fmt.Println(reflect.TypeOf(arg))
	}
}
```

5. Init Directory
```cmd
C:\Example\beeper> go mod init beeper
```

6. Create Cross-Platform Build Bash/Bat file (On Windows I use git bash)
```bash
#!/usr/bin/env bash

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
package_split=(${package//\// })
package_name=${package_split[-1]}
	
platforms=("windows/amd64" "windows/386" "darwin/amd64", "linux/amd64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package_name'-'$GOOS'-'$GOARCH
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done
```

7. Build Binaries by running bash
```bash
$ ./build.bash main.go
```

8. Test the binaries
```bash
$ ./main.go-windows-amd64.exe -img "assets/warning.png"
```