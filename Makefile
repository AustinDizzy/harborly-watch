all: build

build: install_deps
	@go build
	@mkdir -p bin
	@mv harborly-watch ./bin/

clean:
	@for dir in bin log pid ; do \
		sudo rm -rf /opt/harborly-watch/$$dir/* ; \
	done

install: build install_service
	@for dir in bin log pid ; do \
		sudo mkdir -p /opt/harborly-watch/$$dir ; \
	done
	sudo cp ./bin/harborly-watch /opt/harborly-watch/bin/
	@if [ -s config.yaml ] ; \
	then \
		sudo cp ./config.yaml /opt/harborly-watch/bin/ ; \
	fi;

install_deps:
	@go get

install_service:
	@chmod +x ./init.d/harborly-watch
	sudo chown root:root ./init.d/harborly-watch
	sudo cp -p ./init.d/harborly-watch /etc/init.d/harborly-watch
