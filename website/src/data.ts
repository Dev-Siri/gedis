export const VERSION = "2.0.0";

type Platform = "darwin" | "linux";
type Arch = "arm" | "arm64" | "amd64";

export const REPO_LINK = "https://github.com/Dev-Siri/gedis";
const getDownloadLink = (platform: Platform, arch: Arch) =>
  `${REPO_LINK}/releases/download/v${VERSION}/gedis-${VERSION}-${platform}-${arch}.tar.gz`;

export const MAC_VERSIONS = [
  {
    url: getDownloadLink("darwin", "amd64"),
    arch: "x64",
  },
  {
    url: getDownloadLink("darwin", "arm64"),
    arch: "ARM64",
  },
];

export const LINUX_VERSIONS = [
  {
    url: getDownloadLink("linux", "amd64"),
    arch: "x64",
  },
  {
    url: getDownloadLink("linux", "arm"),
    arch: "ARM",
  },
  {
    url: getDownloadLink("linux", "arm64"),
    arch: "ARM64",
  },
];
