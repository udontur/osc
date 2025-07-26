# OSC
OSC is `udontur/ossdmk` written python to make the calculations more precise, the code more maintainable, and reduce the program's complexity.

## Installation (Linux)
1. Add the Nix flake to your ```flake.nix``` input:
```nix
osc.url="github:udontur/ossdmk";
```
2. Add the package to your ```configurations.nix```:
```nix
inputs.osc.packages."${system}".default
```
3. Rebuild your Nix configuration.

## Usage
Enter ```osc``` in the terminal.

