deps:
	@echo "I need sudo access to install the dependencies"
	@if ! hash luarocks; then echo "I need luarocks installed."; false; fi
	sudo luarocks install moonrocks --server=http://rocks.moonscript.org
	sudo moonrocks install yaml
	sudo moonrocks install moonscript

install:
	cp dev.moon ~/bin/dev
