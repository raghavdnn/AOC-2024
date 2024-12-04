{
  description = "Development environment for AOC-2024 project";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells = {
          default = pkgs.mkShell {
            buildInputs = [
              pkgs.go          # Latest Go
            ];

            shellHook = ''
              echo "Welcome to the AOC-2024 development environment!"
              echo "Go: $(go version)"
            '';
          };
        };
      }
    );
}
