########################################################## ENT

ent-binary=entgo.io/ent/cmd/ent
ent-path=./internal/database/ent/schema

ent-get: 
	go get -u $(ent-binary)

ent-init: 
	go run $(ent-binary) init User --target=$(ent-path)

ent-generate:
	go run -mod=readonly $(ent-binary) generate $(ent-path)

##########################################################
