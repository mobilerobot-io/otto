dirs = prov config destroy

echo:
	echo $< $@

prov:
	make -C $< $@

config:
	make -C config $@

destroy:
	make -C config $@
	make -C prov $@

status:
	make -C prov $@
	make -C config $@

clean:
	rm -rf *~ tmp
	make -C prov $@
	make -C config $@

.PHONY: prov
