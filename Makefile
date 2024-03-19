cg:
	cd ./contacts/components && templ generate

cgr:
	cd ./contacts/components && templ generate && cd ../ && go run cmd/main.go

cr:
	cd ./contacts && go run cmd/main.go

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