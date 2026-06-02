.PHONY: gitpush

gitpush:
	@if [ -z "$$(git status --porcelain)" ]; then \
		echo "No changes to commit."; \
	else \
		read -p "Commit message: " msg; \
		git add . && \
		git commit -m "$$msg" && \
		git push origin HEAD; \
	fi