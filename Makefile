# we will put our integration testing in this path
INTEGRATION_TEST_PATH?=./test/it
SRC_PATH?=./internal
ENV_CONFIG?=./config/dev.env
# this command will trigger integration test
# INTEGRATION_TEST_SUITE_PATH is used for run specific test in Golang, if it's not specified
# it will run all tests under ./it directory
test.integration:
	docker-compose --env-file $(ENV_CONFIG) up -d; \
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -run=$(INTEGRATION_TEST_SUITE_PATH);

# this command will trigger integration test with verbose mode
test.integration.debug:
	$(ENV_LOCAL_TEST) \
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -v -run=$(INTEGRATION_TEST_SUITE_PATH)
