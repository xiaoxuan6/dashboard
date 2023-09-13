.PHONY: install,login,dev
install:
	@npm install -g vercel

login:
	@vercel login

dev:
	@vercel dev -y
