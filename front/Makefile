make_version:
	echo "Building: $(VERSION)"
	docker build . -t materkov/shortener-front:$(VERSION)
	docker push materkov/shortener-front:$(VERSION)
