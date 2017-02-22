package main

import (
	"fmt"
	"github.com/ghodss/yaml"
)

func main() {
	str := "## Bitnami MongoDB image version\n## ref: https://hub.docker.com/r/bitnami/mongodb/tags/\n##\nimage: bitnami/mongodb:3.4.1-r0\n\n## Specify a imagePullPolicy\n## 'Always' if imageTag is 'latest', else set to 'IfNotPresent'\n## ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images\n##\n# imagePullPolicy:\n\n## MongoDB admin password\n## ref: https://github.com/bitnami/bitnami-docker-mongodb/blob/master/README.md#setting-the-root-password-on-first-run\n##\n# mongodbRootPassword:\n\n## MongoDB custom user and database\n## ref: https://github.com/bitnami/bitnami-docker-mongodb/blob/master/README.md#creating-a-user-and-database-on-first-run\n##\n# mongodbUsername:\n# mongodbPassword:\n# mongodbDatabase:\n\n## Enable persistence using Persistent Volume Claims\n## ref: http://kubernetes.io/docs/user-guide/persistent-volumes/\n##\npersistence:\n  enabled: true\n  storageClass: generic\n  accessMode: ReadWriteOnce\n  size: 8Gi\n\n## Configure resource requests and limits\n## ref: http://kubernetes.io/docs/user-guide/compute-resources/\n##\nresources:\n  requests:\n    memory: 256Mi\n    cpu: 100m\n"

	var values map[string]interface{}

	_ = yaml.Unmarshal([]byte(str), &values)

	fmt.Printf("%v\n", values)
}
