
local:
	hugo server --minify --theme hugo-book

build:
	hugo --minify --theme hugo-book

deploy: build
	gcloud app deploy

clean:
	rm -rf public resources

update:
	git submodule update --init
