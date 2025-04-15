{ pkgs ? import ../../../nix { } }:
let silcd = (pkgs.callPackage ../../../. { });
in
silcd.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [
    ./broken-silcd.patch
  ];
})
