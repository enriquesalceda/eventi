clean:
	rm -rf ./bin/*
	rm -rf ./deployable/*

build:
	env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o ./bin/specificschedule/bootstrap ./handler/specificschedule/main.go

zip:
	zip -j deployable/specificschedule.zip ./bin/specificschedule/bootstrap

deploy:
	sls deploy

build.zip.deploy:
	make build && make zip && sls deploy

clean.deploy:
	make clean && make builds.zips.deploy