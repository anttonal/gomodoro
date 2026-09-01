{
  description = "gomodoro - terminal pomodoro timer";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-26.05";

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
    in {
      packages.${system}.default = pkgs.buildGoModule {
        pname = "gomodoro";
        version = "0.1.0";
        src = ./.;
        # Standard library only, so there are no vendored deps to hash.
        vendorHash = null;
      };

      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          go
          gopls
          golangci-lint
          delve
        ];

        shellHook = ''
          export GOTOOLCHAIN=local
        '';
      };
    };
}
