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

	-mkdir content/$(path)

	@echo '---'           >> content/$(path)/index.md
	@echo 'title: Fix Me' >> content/$(path)/index.md
	@echo 'description:'  >> content/$(path)/index.md
	@echo 'draft: true'   >> content/$(path)/index.md
	@echo "date: "$(NOW)""  >> content/$(path)/index.md
	@echo '---'           >> content/$(path)/index.md

.PHONY: save
save: ## save posts
	@git checkout write
	@git add .
	@git commit -m ":memo:save on ["$(NOW)"]"
