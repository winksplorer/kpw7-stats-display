NAME = kpw7-stats-display

GO ?= go
TAR ?= tar
PREFIX ?= /usr


.PHONY: backend all sea clean install

all: backend sea

clean:
	rm -rf $(NAME)-server $(NAME) $(NAME).tar.xz $(NAME).tar.xz.o frontend-dist

backend:
	@echo "     GO (${GO}) backend -> $(NAME)-server"
	@cd backend && $(GO) mod tidy && $(GO) build -o ../$(NAME)-server -ldflags="-s -w"

# compress frontend+backend, then embed it into a binary combined with compiled SEA
sea:
	@echo "     XZ ($(TAR)) $(NAME)-server + frontend -> $(NAME).tar.xz"
	@$(TAR) -cJf $(NAME).tar.xz $(NAME)-server frontend
	@echo "     LD ($(LD)) $(NAME).tar.xz -> $(NAME).tar.xz.o"
	@$(LD) -r -b binary -o $(NAME).tar.xz.o $(NAME).tar.xz
	@echo "     CC ($(CC)) sea.c + $(NAME).tar.xz.o -> $(NAME)"
	@$(CC) -o $(NAME) sea.c $(NAME).tar.xz.o -larchive

install:
	install -m 755 $(NAME) $(PREFIX)/bin