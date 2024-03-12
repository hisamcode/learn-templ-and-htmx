g:
	cd ./hypermedia.system/components && templ generate
	
gr:
	cd ./hypermedia.system/components && templ generate && cd ../ && go run cmd/main.go

r:
	cd ./hypermedia.system && go run cmd/main.go

hs:
	cd ./hypermedia.system && air

templ:
	cd ./templ && air