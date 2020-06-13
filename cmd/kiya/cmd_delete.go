package main

import (
	"fmt"
	"log"

	cloudkms "cloud.google.com/go/kms/apiv1"
	cloudstore "cloud.google.com/go/storage"
	"github.com/emicklei/tre"

	"github.com/kramphub/kiya"
)

// commandDelete deletes a stored key
func commandDelete(kmsService *cloudkms.KeyManagementClient, storageService *cloudstore.Client, target kiya.Profile, key string) {
	_, err := kiya.GetValueByKey(kmsService, storageService, key, target)
	if err != nil {
		log.Fatal(tre.New(err, "delete failed", "key", key, "err", err))
	}
	if promptForYes(fmt.Sprintf("Are you sure to delete [%s] from [%s] (y/N)? ", key, target.Label)) {
		if err := kiya.DeleteSecret(storageService, target, key); err != nil {
			fmt.Printf("failed to delete [%s] from [%s] because [%v]\n", key, target.Label, err)
		} else {
			fmt.Printf("Successfully deleted [%s] from [%s]\n", key, target.Label)
		}
	} else {
		log.Fatalln("delete aborted")
	}
}