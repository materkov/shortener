make_version:
	echo "Building: $(VERSION)"
	docker build . -t materkov/shortener-redirecter:$(VERSION)
	docker push materkov/shortener-redirecter:$(VERSION)
