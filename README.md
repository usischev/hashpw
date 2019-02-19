# hashpw

Command-line utility to hash a password with PBKDF2-SHA512 in a way compatible with OpenLDAP

## Usage

    ./hashpw

Or:

    echo My_Super_Password | ./hashpw

One password is read from standard input. Password prompt is displayed if input is not pipeline.
One line is written to output formatted for OpenLDAP.

## Building

Making size-optimized builds for macOS and Linux (64-bit):

    GOOS=darwin GOARCH=amd64 go build -o hashpw_macos -ldflags="-s -w"
    GOOS=linux GOARCH=amd64 go build -o hashpw_linux -ldflags="-s -w"

Optionally, use UPX to reduce size of executables:

    upx --brute hashpw_macos hashpw_linux
