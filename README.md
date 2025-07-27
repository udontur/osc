# OSC
OSC is `udontur/ossdmk` written in python to provide more accurate calculations, maintainable code, and reduce the program's complexity.

It is a small terminal utility that calculates your OSSD mark. 

## Installation (Linux or MacOS)
> [!IMPORTANT]
> You must have Python3 installed on your system

1. Add the Nix flake to your ```flake.nix``` input:
```nix
osc.url="github:udontur/osc";
```
2. Add the package to your ```configurations.nix```:
```nix
inputs.osc.packages."${system}".default
```
3. Rebuild your Nix configuration.

## Usage
Enter ```osc``` in the terminal.
> [!NOTE]
> It can handle a lot of decimal points, but it will be rounded to 3 decimal places at the end. 
