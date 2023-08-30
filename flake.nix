{
  description = "terraform-provider-bria";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = nixpkgs.legacyPackages.${system};
      buildDeps = with pkgs; [
        go
        gotools
        gnumake
        tfplugindocs
      ];
      buildInputs = with pkgs;
        buildDeps
        ++ [
          protobuf
          protoc-gen-go
          protoc-gen-go-grpc
          terraform
        ];
    in
      with pkgs; rec {
        packages = {
          terraform-provider-bria = buildGoModule rec {
            pname = "terraform-provider-bria";
            version = "0.1.0";

            src = ./.;

            vendorSha256 = null;
          };
        };

        packages.default = packages.terraform-provider-bria;

        devShells.default = mkShell {
          inherit buildInputs;
          packages = [alejandra vendir gopls];
        };

        formatter = alejandra;
      });
}
