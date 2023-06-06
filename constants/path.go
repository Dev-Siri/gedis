package constants

import "fmt"

const CacheFolder string = ".cache";
const BackupFolder string = "backups"
var MapJsonPath string = fmt.Sprintf("%s/map.json", CacheFolder)