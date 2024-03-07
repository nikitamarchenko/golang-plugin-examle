build-plugins:
	$(foreach file, $(wildcard ./plugins/*), $(MAKE) -C ${file};)
