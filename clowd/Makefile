dirs = prov config destroy

prov:
	make -C prov $@

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
