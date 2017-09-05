### kibanator-server

Proxying server to serve [kibanator](https://github.com/mgurov/kibanator)

### Build 

At the kibanator folder: `REACT_APP_API_PATH="/api" yarn build`

````bash
mv <kibanator>/build/ ./ui
go-bindata-assetfs ui/...
mv bindata_assetfs.go main/
go build -o kibanator-server main/*.go
````
