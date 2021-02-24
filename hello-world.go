package main

import "fmt"

func main() {
	fmt.Println("Hello from Docker on GitHub Container Registry!\nThis message shows that your installation appears to be working correctly.\n\nAs Docker Inc introduced rate-limiting in https://www.docker.com/increase-rate-limits\nwe simply need our own hello-world image on GitHub Container Registry.\n\nTo generate this message, Docker took the following steps:\n 1. The Docker client contacted the Docker daemon.\n 2. The Docker daemon pulled this \"hello-world\" image from the GitHub Container Registry.\n    (amd64)\n 3. The Docker daemon created a new container from that image which runs the\n    executable that produces the output you are currently reading.\n 4. The Docker daemon streamed that output to the Docker client, which sent it\n    to your terminal.\n\n")
}
