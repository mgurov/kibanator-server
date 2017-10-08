.PHONY: mvui ui rebind run
mvui:
	rm -Rf ./ui
	mv ../kibanator/build/ ./ui
rebind: 
	go-bindata-assetfs ui/...
	mv bindata_assetfs.go main/
ui: mvui rebind
run: 
	go run main/*.go