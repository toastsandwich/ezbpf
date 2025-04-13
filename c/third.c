#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <linux/udp.h>
#include <bpf/bpf_helpers.h>

struct  data_t {
	__be32 src_addr;
__u64 src__port;	
};			

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__type(key, __u32);
	__type(value, __u64);
	__uint(max_entries, 1024);
} my_map .SEC("maps");

int add(int a,int b) {
	return a + b;
}


int all_pass() {
	int a = 10;

if (a == 10) {
	return 1;
}

return 0;
}

