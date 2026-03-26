{
  description = "Find dirty git repo's" ;

  inputs.nixpkgs.url = "nixpkgs/nixos-24.05";

  outputs = { self, nixpkgs }:
    let
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });
      version = builtins.replaceStrings ["\n"] [""] (builtins.readFile ./src/VERSION);
    in
    {
      nixosModules.default = import ./module.nix self;

      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          dirty-repo-scanner = pkgs.callPackage ./package.nix { inherit version; };
        });

      defaultPackage = forAllSystems (system: self.packages.${system}.dirty-repo-scanner);

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
            ];
          };
        });
    };
}

