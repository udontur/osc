{
  description = "github:udontur/osc Nix flake";
  inputs.nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  
  outputs = { self, nixpkgs, ... }:
    let
      supportedSystems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
    in{
      packages = forAllSystems(system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in{
        default =
          pkgs.stdenv.mkDerivation rec {
            pname = "osc";
            version = "1.0";
            src = self;

            installPhase = ''
              runHook preInstall
              
              mkdir -p $out/bin
              install -Dm755 ./osc $out/bin/osc

              runHook postInstall
            '';
            
            meta = {
              homepage = "https://github.com/udontur/osc";
              mainProgram = "osc";
              platforms = pkgs.lib.platforms.all;
            };
          };
        }
      );
    };
}
