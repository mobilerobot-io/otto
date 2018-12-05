test stage prod: .PHONY
	$(MAKE) -C $< $@

hosts: hosts-*
	echo "[all]" > hosts
	cat hosts-* >> hosts
	echo "" >> hosts
	echo "[nginx]" >> hosts
	cat hosts-w* >> hosts
	echo "" >> hosts
	echo "[otto]" >> hosts
	cat hosts-otto >> hosts
	echo "" >> hosts

clean:
	rm *~
