{ lib, buildGoModule, go, version ? "0.2.0" }:
buildGoModule rec {
  pname = "dirty-repo-scanner";

  inherit version;

  src = ./.;

  subPackages = [ "src" ];

  postInstall = ''
    mv $out/bin/src $out/bin/drs
  '';

  doCheck = false;

  vendorHash = "sha256-yn2hTbEkYLJrAxjbCDpW2V8U4EPcOyhZTQdFNvcJRGs=";

  meta = with lib; {
    description = ''
      Find dirty repos
    '';
    homepage = "https://github.com/mipmip/dirty-repo-scanner";
    mainProgram = "drs";
    license = licenses.bsd2;
  };

}
