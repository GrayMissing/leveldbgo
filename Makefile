leveldb:
	mkdir build
	cd build && cmake -DCMAKE_BUILD_TYPE=Release -DBUILD_SHARED_LIBS=ON -DLEVELDB_BUILD_TESTS=OFF -DLEVELDB_BUILD_BENCHMARKS=OFF ../leveldb
	mkdir dist
	cmake --build ./build
	cmake --install ./build --prefix ./dist

.PHONY : leveldb

clean:
	rm -r build
	rm -r dist

.PHONY : clean