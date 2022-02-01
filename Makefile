build:
	DOCKER_BUILDKIT=1 docker build -t halmakidon/httpecho:1.0.0 .

run:
	docker run -it --rm -p 8080:8080 halmakidon/httpecho:1.0.0