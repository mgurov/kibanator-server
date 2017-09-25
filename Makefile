.PHONY: ui
ui:
	rm -Rf ./ui
	mv ../kibanator/build/ ./ui
	go-bindata-assetfs ui/...
	mv bindata_assetfs.go main/