#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <linux/udp.h>
#include <bpf/bpf_helpers.h>

struct  foo {
	int bar;
__u32 cat;	
};			

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__type(key, __u32);
	__type(value, __u32);
	__uint(max_entries, 1024);
} mm .SEC("maps");

struct foo Foo() {
	struct foo f = {9,10};
return f;
}


int new() {
	int i = 10;
long j = i + 10;
__u32 k = j * j + 10;
struct foo g = {0,0};
struct foo f = Foo();
bpf_map_update_elem(m,k,j,BPF_ANY);
return f;
}


int temp() {
	return 190;
}

