# rMviewer

A user-friendly interface to the installation and usage of rMview for reMarkable.

### Installation

- Go to the [releases page](https://github.com/rhobro/rMviewer/releases) and download the latest binary of `rmviewer` in
  the assets section. Please choose the "amd64" version if you are running an Intel processor based Mac or choose the
  "arm64" version if you are running on the M1 chip.
- Note that when you run any following commands containing "rmviewer", you will have to use the full name of the file
  which you downloaded.
- Open the Terminal app on Mac OS and navigate to the downloads folder using `cd Downloads` (this is the default unless
  you have changed it).
- Allow the file to be executed by running `chmod +x rmviewer`

### Usage

- Run the executable using `./rmviewer`

This should install all the necessary dependencies such as Homebrew, Python 3 and the relevant sub-dependencies before
installing rMview itself. The executable will automatically start rMview after it has been installed. The same
executable can be run in the future by clicking on the file in Finder, but will only install dependencies again if they
have been removed since.

Otherwise, it is now possible to start rMview by opening Terminal anywhere and running `rmview`.