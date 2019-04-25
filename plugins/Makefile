plugins := $(wildcard $(dir $(wildcard */*.so)))

# $(plugins): $(plugins)/$<.so
#	$(MAKE) -C $< $@

$(plugins):
	@echo $@

status:
	@echo $(plugins)

pi: $(plugins)
	$(MAKE) -C $< $@

.PHONY: status $(plugins)
