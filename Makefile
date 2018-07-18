ifeq ($(OS), Windows_NT)
	OSName = Windows
else
	OSName = $(shell echo $(shell uname -s))
	OSArch = $(shell echo $(shell uname -m))
	ifeq ($(OSArch), x86_64)
		OSArch = amd64
	else ifeq ($(OSArch), armv7l)
		OSArch = armv7l
	endif
endif

depend: $(OSName)

$(OSName):
	$(RM) -r driver/include
	$(RM) -r driver/$@-lib
	$(RM) -r temp-driver
	cd 3rdParty/mongo-c-driver && \
	if [ -d dist ]; then \
		rm -rf dist; \
	fi && \
	mkdir dist && \
	cd dist && \
	cmake -DCMAKE_INSTALL_PREFIX=../../../temp-driver .. && \
	make && make install
	cp -R temp-driver/lib driver/$@-lib
	cp -R temp-driver/include driver/include
	cd driver && rm -rf $@-lib/*.so $@-lib/*.so.* $@-lib/*.dylib $@-lib/cmake $@-lib/pkgconfig
	$(RM) -r 3rdParty/mongo-c-driver/dist

lint:
	gometalinter.v2 ./...

test:
	go test ./...

