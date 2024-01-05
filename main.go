package main

import (
	"os"
	"fmt"
	host_client "github.com/matehaxor03/holistic_host_client/host_client"

)

func main(){
	host_client, host_client_errors := host_client.NewHostClient()
	if host_client_errors != nil {
		fmt.Println(fmt.Errorf("%s", host_client_errors))
		os.Exit(1)
	}

	ramdisk, ramdisk_errors := host_client.Ramdisk("ramdisk", uint64(2048*1000))
	if ramdisk_errors != nil {
		fmt.Println(fmt.Errorf("%s", ramdisk_errors))
		os.Exit(1)
	}

	ramdisk_create_errors := ramdisk.Create()
	if ramdisk_create_errors != nil {
		fmt.Println(fmt.Errorf("%s", ramdisk_create_errors))
		os.Exit(1)
	}

	fmt.Println("install successfull")
	os.Exit(0)
}