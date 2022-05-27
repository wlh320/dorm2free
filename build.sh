#!/bin/bash
# go build -tags=legacy_appindicator
go build
strip portguard-systray2
upx --lzma --best portguard-systray2
