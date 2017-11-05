package setup

const (
	VERSION = "1.0.3"
	BASE_DOCKER_TAG = "0.3.2"
	BASE_DOCKER_NAME = "baseos"
	MIN_DOCKER_VER = "17.06.2"
	DOCKER_IMG_PREFIX = "hyperledger/fabric-"
)

var IMAGES = [...]string {"peer", "orderer", "couchdb", "ccenv", "javaenv", "kafka", "zookeeper", "tools", "ca"}
var SUPPORTED_ARCHS = [...]string {"amd64", "ppc64le", "s390x"}
var SUPPORTED_OS = [...]string {"darwin", "linux", "windows"}
var DOWNLOADS = [...]string { "https://nexus.hyperledger.org/content/repositories/releases/org/hyperledger/fabric/hyperledger-fabric/" + GetBinArch() + "-" + VERSION + "/" + "hyperledger-fabric-" + GetBinArch() + "-" + VERSION + ".tar.gz",
//TODO Invalid Location from get-byfn.sh	"https://nexus.hyperledger.org/content/repositories/releases/org/hyperledger/fabric/examples/hyperledger-fabric-byfn-" + VERSION + ".tar.gz",
}
