########################################################## ENTs

ent-binary=entgo.io/ent/cmd/ent
ent-path=./internal/database/ent/schema

ent-get: 
	go get -u $(ent-binary)

ent-init: ent-get
	go run $(ent-binary) init Book Author --target=$(ent-path)

ent-generate: ent-init
	go run -mod=readonly $(ent-binary) generate $(ent-path)

########################################################## TESTs

test: ent-generate
	go test -v ./... -covermode=atomic -coverprofile=coverage.out
