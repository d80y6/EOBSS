# TelcoFlow Project Makefile

SERVICES := iam-service crm-service catalog-service order-service billing-service inventory-service assurance-service provisioning-service
APPS := admin-portal

.PHONY: build test docker-build clean

build:
	@for service in $(SERVICES); do \
		echo "Building $$service..."; \
		cd services/$$service && go build ./... && cd ../..; \
	done

test:
	@for service in $(SERVICES); do \
		echo "Testing $$service..."; \
		cd services/$$service && go test ./... && cd ../..; \
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
