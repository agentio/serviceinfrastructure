
local:
	hugo server --minify --theme hugo-book

build:
	hugo --minify --theme hugo-book

deploy:
	gcloud app deploy

clean:
	rm -rf public resources
