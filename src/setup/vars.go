package setup

const (
	FABRICTAG = ""
	VERSION = "1.0.3"
	MIN_DOCKER_VER = "17.06.2"
	DOCKER_IMG_PREFIX = "hyperledger/fabric-"
	BIN = "bin"
	BIN_DL_ROOT = "https://nexus.hyperledger.org/content/repositories/releases/org/hyperledger/fabric/hyperledger-fabric/"
)

var IMAGES = [...]string {"peer", "orderer", "couchdb", "ccenv", "javaenv", "kafka", "zookeeper", "tools", "ca"}
var SUPPORTED_ARCHS = [...]string {"amd64", "ppc64le", "s390x"}
var SUPPORTED_OS = [...]string {"darwin", "linux", "windows"}
