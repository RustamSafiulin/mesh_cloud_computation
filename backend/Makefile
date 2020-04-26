all: clean build

clean:
	make -C account_service/ clean
	make -C job_service/ clean

build:
	make -C account_service/ build
	make -C job_service/ build