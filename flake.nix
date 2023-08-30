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
      pkgs = import nixpkgs {inherit system;};
      nativeBuildInputs = with pkgs; [
        gnumake
        tfplugindocs
        protobuf
        protoc-gen-go
        protoc-gen-go-grpc
      ];
      buildInputs = with pkgs; [
        go
        gotools
      ];
      terraform-provider-bria = pkgs.buildGoModule {
        pname = "terraform-provider-bria";
        version = "0.1.0";

        src = ./.;

        vendorSha256 = null;
      };
    in
      with pkgs; {
        packages = {
          inherit terraform-provider-bria;
          default = terraform-provider-bria;
        };

        devShells.default = mkShell {
          inherit buildInputs nativeBuildInputs;
          packages = [
            alejandra
            vendir
            gopls
            terraform
          ];
        };

        formatter = alejandra;
      });
}
