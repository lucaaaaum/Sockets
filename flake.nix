{
  description = "Sockets";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
        packages = with pkgs; [
          go
          gopls
          gotools
          go-tools
        ];
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "sockets";
          version = "0.1.0";
          src = ./.;
          vendorHash = null;
          nativeBuildInputs = packages;
          subpackages = [ "cmd/server" "cmd/client" ];
        };
        devShells = {
          default = pkgs.mkShell {
            buildInputs = packages;
          };
        };
      }
    );
}
