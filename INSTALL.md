
# Installation

*pkgasset* is a command line program run from a shell like Bash. You can find compiled
version in the [releases](https://github.com/caltechlibrary/pkgasset/releases/latest). 
Download the zip file and unzip it. The filename is in the form of `pkgasset-VERSION_NO-release.zip`.
Inside the zip file look for the directory that matches your computer and copy that someplace
defined in your path (e.g. $HOME/bin if you're running things on Unix/Linux/Mac OS X). 

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows
(amd64) and Rapsberry Pi (ARM7)

## Mac OS X

1. Go to [github.com/caltechlibrary/pkgasset/releases/latest](https://github.com/caltechlibrary/pkgasset/releases/latest)
2. Click on the green "pkgasset-VERSION_NO-release.zip" link and download
3. Open a terminal and type `cd ~/Downloads/` then unzip the file `unzip pkgasset-VERSION_NO-release.zip` and `cd dist/macosx-amd64/`
4. Copy the *pkgasset* to a "bin" directory in your path.  For example, type `sudo cp pkgasset /usr/local/bin`
5. Test by typing `pkgasset -h`

## Windows

1. Go to [github.com/caltechlibrary/pkgasset/releases/latest](https://github.com/caltechlibrary/pkgasset/releases/latest)
2. Click on the green "pkgasset-release.zip" link and download
3. Open the file manager find the downloaded file and unzip it (e.g. pkgasset-release.zip)
4. Look in the unziped folder and find dist/windows-amd64/pkgasset.exe
5. Drag (or copy) the *pkgasset.exe* to a "bin" directory in your path (a good option is C\Users\username\bin)
6. Open Bash and and test by typing `pkgasset -h`
7. If it doesn't work type `echo $PATH` and copy *pkgasset.exe* to one of the directories listed

## Linux

1. Go to [github.com/caltechlibrary/pkgasset/releases/latest](https://github.com/caltechlibrary/pkgasset/releases/latest)
2. Click on the green "pkgasset-release.zip" link and download
3. find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/pkgasset-release.zip)
4. In the unziped directory and find for dist/linux-amd64/pkgasset
5. copy the *pkgasset* to a "bin" directory (e.g. cp ~/Downloads/pkgasset-release/dist/linux-amd64/pkgasset ~/bin/)
6. From the shell prompt run `pkgasset -h`

## Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Go to [github.com/caltechlibrary/pkgasset/releases/latest](https://github.com/caltechlibrary/pkgasset/releases/latest)
2. Click on the green "pkgasset-release.zip" link and download
3. find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/pkgasset-release.zip)
4. In the unziped directory and find for dist/raspberrypi-arm7/pkgasset
5. copy the *pkgasset* to a "bin" directory (e.g. cp ~/Downloads/pkgasset-release/dist/raspberrypi-arm7/pkgasset ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
6. From the shell prompt run `pkgasset -h`

