make_version:
	echo "Building: $(VERSION)"
	docker build . -t materkov/shortener-analytics:$(VERSION)
	docker push materkov/shortener-analytics:$(VERSION)
