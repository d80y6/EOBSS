# TelcoFlow Project Makefile

SERVICES := iam-service crm-service catalog-service order-service billing-service inventory-service assurance-service provisioning-service mediation-service partner-service incident-service
APPS := admin-portal self-care-portal

.PHONY: build test docker-build clean

build:
	@for service in $(SERVICES); do \
		echo "Building $$service..."; \
		if [ -d "services/$$service" ]; then \
			cd services/$$service && go build ./... && cd ../..; \
		fi \
	done

test:
	@for service in $(SERVICES); do \
		echo "Testing $$service..."; \
		if [ -d "services/$$service" ]; then \
			cd services/$$service && go test ./... && cd ../..; \
		fi \
	done

docker-build:
	@for service in $(SERVICES); do \
		docker build -t telcoflow/$$service:latest -f services/$$service/Dockerfile . ; \
	done
	@for app in $(APPS); do \
		docker build -t telcoflow/$$app:latest -f apps/$$app/Dockerfile . ; \
	done

clean:
	@find . -name "*-server" -type f -delete
	@find . -name "venv" -type d -exec rm -rf {} +
