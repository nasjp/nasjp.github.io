# variables
NOW := "`date +'%y/%m/%d-%H:%M:%S'`"
_NOW := "`date +'%y%m%d%H%M%S'`"

# commands
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: serve
serve: ## run hugo server
	@hugo server -D

.PHONY: new
new: ## gen new post
	@hugo new post/post$(_NOW).md

.PHONY: save
save: ## save posts
	@git checkout write
	@git add .
	@git commit -m ":memo:save on ["$(NOW)"]"

.PHONY: delploy
deploy: ## deploy posts
	@git checkout write
	@git push origin write
