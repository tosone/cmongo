all:
	echo $< $? $@

ifeq ($(OS),Windows_NT)
  OSName = windows
  Suffix = .exe
else
  OSName = $(shell echo $(shell uname -s))
endif

depend: ${OSName}

${OSName}:
	-rm -rf driver/include
	-rm -rf driver/$@-lib
	-rm -rf temp-driver
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
	-rm -rf 3rdParty/mongo-c-driver/dist
