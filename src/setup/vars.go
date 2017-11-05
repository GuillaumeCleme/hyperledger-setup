package setup

const (
	FABRICTAG = ""
	VERSION = "1.0.3"
	MIN_DOCKER_VER = "17.06.2"
	DOCKER_IMG_PREFIX = "hyperledger/fabric-"
	BIN = "bin"
)

var IMAGES = [...]string {"peer", "orderer", "couchdb", "ccenv", "javaenv", "kafka", "zookeeper", "tools", "ca"}
var SUPPORTED_ARCHS = [...]string {"x86_64", "ppc64le", "s390x"}
