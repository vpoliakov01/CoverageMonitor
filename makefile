.PHONY: server client

server:
	cd back_end && go build

server-run: server
	cd back_end && ./back_end

client:
	cd front_end && npx webpack

nginx:
	cp nginx.conf /etc/nginx/nginx.conf
	service nginx start

app: nginx server server-run

docker:
	docker build -t vp-CoverageMonitor . && \
	docker run -p 8321:8321 vp-CoverageMonitor &
