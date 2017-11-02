package setup

const (
	FABRICTAG = ""
	VERSION = "1.0.3"
	MIN_DOCKER_VER = "17.06.2"
)

var IMAGES = [...]string {"peer", "orderer", "couchdb", "ccenv", "javaenv", "kafka", "zookeeper", "tools"}
