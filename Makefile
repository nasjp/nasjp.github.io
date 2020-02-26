# variables
NOW := "$$(date +'%Y-%m-%dT%H:%M:%S+09:00')"

# commands
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: develop
develop: ## run gatsby server locally
	@gatsby develop

.PHONY: build
build: ## build gatsby
	@gatsby build

.PHONY: new
new: ## gen new post ($ make new path=test-hoge)
	@$(if $(path),, \
		echo "Please set 'path'"; exit 1\
	)

	@mkdir -p content/blog/$(path)

	@echo '---'            >> content/blog/$(path)/index.md
	@echo 'title: Fix Me'  >> content/blog/$(path)/index.md
	@echo "date: "$(NOW)"" >> content/blog/$(path)/index.md
	@echo 'draft: true'    >> content/blog/$(path)/index.md
	@echo 'tags: ["***"]'  >> content/blog/$(path)/index.md
	@echo '---'            >> content/blog/$(path)/index.md

.PHONY: save
save: ## save posts
	@git checkout write
	@git add .
	@git commit -m ":memo:save on ["$(NOW)"]"

.PHONY: deploy
deploy: ## deploy posts
	@git checkout write
	@git push origin write
