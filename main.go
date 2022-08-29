package main

import(
	"go_redis/tcp"
)

func main() {
    tcp.ListenAndServe(":8000")
}