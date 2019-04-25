plugins := $(wildcard $(dir $(wildcard */*.so)))

all: $(plugins)

$(plugins):
	$(MAKE) -C $< $@

status:
	@echo $(plugins)

pi: $(plugins)
	$(MAKE) -C $< $@

.PHONY: status pi all $(plugins)
