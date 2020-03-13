VER := $(shell git describe --tags --always --dirty || echo unknown)
IMG := kentik-proto-build:$(VER)

GENERATED :=		            \
	darknet/darknet.pb.go	    \
	darknet/darknet.pb-c.c	    \
	darknet/darknet.pb-c.h

build: img
	$(eval CID := $(shell docker run -di $(IMG) bash))
	-$(MAKE) all CID=$(CID)
	docker kill $(CID)
	docker rm   $(CID)

img:
	docker build -t $(IMG) .

all: $(GENERATED)

clean:
	-docker rmi $(IMG)
	$(RM) $(GENERATED)

# ----------------------------------------------------------------

DOEX = docker exec $(CID)
DOCP = docker cp $(CID):/proto/$@ $@

%.pb.go: %.proto
	$(DOEX) protoc --go_out=plugins=grpc:. $<
	$(DOCP)

%.pb-c.c: %.proto
	$(DOEX) bash -c "cd $(dir $<) && protoc-c --c_out=. $(notdir $<)"
	$(DOEX) sed -i "1i // +build ignore\n" $@
	$(DOCP)

%.pb-c.h: %.pb-c.c
	$(DOCP)

%.pb.cc: %.proto
	$(DOEX) bash -c "cd $(dir $<) && protoc --cpp_out=. $(notdir $<)"
	$(DOEX) sed -i "1i // +build ignore\n" $@
	$(DOCP)

%.pb.h: %.pb.cc
	$(DOCP)

.PHONY: all build clean img
