{
  description = "github:udontur/osc Nix flake";
  inputs.nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  outputs =
    { self, nixpkgs }:
    let
      supportedSystems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
    in
    {
      packages = forAllSystems (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
          my-script = (pkgs.writeScriptBin "osc" (builtins.readFile ./osc)).overrideAttrs (old: {
            buildCommand = "${old.buildCommand}\n patchShebangs $out";
          });
          src = ./.;
        in
        {
          default = pkgs.symlinkJoin {
            name = "osc";
            paths =
              [ my-script ] ++ [
                pkgs.python312
              ];
            buildInputs = [ pkgs.makeWrapper ];
            postBuild = ''
              cp ${src}/main.py $out/bin/main.py
              wrapProgram $out/bin/osc --prefix PATH : $out/bin
            '';
          };
        }
      );
    };
}
