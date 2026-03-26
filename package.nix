{ lib, buildGoModule, go }:
buildGoModule rec {
  pname = "dirty-repo-scanner";

  version = "1.0";

  src = ./.;

  doCheck = false;

  vendorHash = "sha256-KBu77tQfZjZsAcUatXZj+sHa+5uUNN5PuFaSk1rzIkQ=";

  meta = with lib; {
    description = ''
      Find dirty repos
    '';
    homepage = "https://github.com/mipmip/dirty-repo-scanner";
    license = licenses.bsd2;
  };

}
