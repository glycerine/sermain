all:
	go build  -gcflags "-N -l"
	go install