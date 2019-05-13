plugins := echo dork periph static store wally wsgi

all: $(plugins)

$(plugins):
	$(MAKE) -C $< $@

%.so : %.go
	$(MAKE) -C $<

status:
	@echo $(plugins)

pi: $(plugins)
	$(MAKE) -C $< $@

.PHONY: status pi all $(plugins)
