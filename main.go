package main

import(
	"go_redis/tcp"
)

func main() {
    tcp.SimpleListenAndServe(":8000")
}