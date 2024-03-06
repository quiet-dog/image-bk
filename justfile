run:
    swag init  && go build main.go && ./main
dev:
    cd web && pnpm dev
web:
    cd web && pnpm build