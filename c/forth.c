#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <linux/udp.h>
#include <bpf/bpf_helpers.h>

struct  pkt_info_t {
	__be32 src_addr;
__u64 src_port;
__u64 payload;	
};			

struct pkt_info_t sample_pkt_info() {
	struct pkt_info_t pkt = {};
return pkt;
}


void main() {
	struct pkt_info_t sample = sample_pkt_info();;
}

