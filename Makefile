all:
	godep go build -i -o food

clean:
	rm -fr food
