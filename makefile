GO_RUN=go run
PRISMA=github.com/steebchen/prisma-client-go

generate:
	$(GO_RUN) $(PRISMA) generate

db-push:
	$(GO_RUN) $(PRISMA) db push

db-pull:
	$(GO_RUN) $(PRISMA) db pull

migrate-dev:
	$(GO_RUN) $(PRISMA) migrate dev

migrate-deploy:
	$(GO_RUN) $(PRISMA) migrate deploy

run:
	go run main.go

clean:
	rm -rf prisma-client/

init:
	go mod tidy

help:
	@echo "Comandos disponíveis:"
	@echo "  generate       - Regera o cliente do Prisma"
	@echo "  db-push        - Sincroniza o banco com o esquema (dev)"
	@echo "  db-pull        - Cria o schema Prisma a partir do banco existente"
	@echo "  migrate-dev    - Cria uma migração localmente para desenvolvimento"
	@echo "  migrate-deploy - Sincroniza o banco de produção com as migrações"
	@echo "  run            - Roda o projeto"
	@echo "  clean          - Remove arquivos gerados"
