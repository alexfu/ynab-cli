{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    let
      mkOutputs = system:
        let
          pkgs = import nixpkgs { inherit system; };
        in
          {
            devShells.default = with pkgs; mkShell {
              packages = [
                # list packages here
              ];
            };
          };
    in
      flake-utils.lib.eachDefaultSystem mkOutputs;
}
