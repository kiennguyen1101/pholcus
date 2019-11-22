go-bindata-assetfs views/...
sed -i 's/package main/package web/' bindata_assetfs.go

sed -i 's|"strings"|"strings";"os"|' bindata_assetfs.go
# sed -i 's/views_index_html/viewsIndexHtmlBytes/g' bindata_assetfs.go
