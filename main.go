package main

import (
	"os"
	host_client "github.com/matehaxor03/holistic_host_client/host_client"

)

func main(){
	host_client, host_client_errors := host_client.NewHostClient()
	if host_client_errors != nil {
		os.Exit(1)
	}

	
	ramdisk, ramdisk_errors := host_client.Ramdisk("ramdisk", uint64(2048*1000))
	if ramdisk_errors != nil {
		os.Exit(1)
	}

	ramdisk_create_errors := ramdisk.Create()
	if ramdisk_create_errors != nil {
		os.Exit(1)
	}

	os.Exit(0)
}