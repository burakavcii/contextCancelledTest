run:
	go test -count 1000 -v ./httpClients -run $(test)

runAll:
	make run test=".*"